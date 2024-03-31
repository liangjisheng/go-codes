package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoTest"
	app.Usage = "hello world"
	app.Version = "1.0.0"

	var language string

	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:  "port, p",
			Value: 8080,
			Usage: "listening port",
		},
		&cli.StringFlag{
			Name:  "lang, l",
			Value: "en",
			Usage: "read from `FILE`",
			// 默认使用第一个
			EnvVars:     []string{"LEGACY_COMPAT_LANG", "APP_LANG", "LANG"},
			Destination: &language,
		},
	}

	// 修改 help 和 version 这2个flag默认形式
	//cli.HelpFlag = &cli.BoolFlag{
	//	Name: "help, h",
	//	Usage: "Help!Help!",
	//}
	//
	//cli.VersionFlag = &cli.BoolFlag{
	//	Name: "print-version, v",
	//	Usage: "print version",
	//}

	// 添加 command
	app.Commands = []*cli.Command{
		{
			Name:     "add",
			Aliases:  []string{"a"},
			Usage:    "calc 1+1",
			Category: "arithmetic",
			Action: func(c *cli.Context) error {
				fmt.Println("1 + 1=", 1+1)
				return nil
			},
		},
		{
			Name:     "sub",
			Aliases:  []string{"s"},
			Usage:    "calc 5-3",
			Category: "arithmetic",
			Action: func(c *cli.Context) error {
				fmt.Println("5 - 3 =", 5-3)
				return nil
			},
		},
		{
			Name:     "db",
			Usage:    "database operations",
			Category: "database",
			Subcommands: []*cli.Command{
				{
					Name:  "insert",
					Usage: "insert data",
					Action: func(c *cli.Context) error {
						fmt.Println("insert subcommand")
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "delete data",
					Action: func(c *cli.Context) error {
						fmt.Println("delete subcommand")
						return nil
					},
				},
			},
		},
	}

	// 如果你想在command执行前后执行后完成一些操作，可以指定app.Before/app.After这两个字段
	app.Before = func(c *cli.Context) error {
		fmt.Println("app before")
		return nil
	}
	app.After = func(c *cli.Context) error {
		fmt.Println("app after")
		return nil
	}

	app.Action = func(c *cli.Context) error {
		//fmt.Println("BOOM")
		fmt.Println(c.String("lang"), c.Int("port"))
		//fmt.Println(language)
		name := "someone"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if language == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}

		return nil
	}

	// 可以对 flag 和 command 进行排序
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
