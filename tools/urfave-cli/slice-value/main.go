package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	// set args for examples sake
	os.Args = []string{"multi_values",
		"--stringSlice", "parsed1", "--stringSlice", "parsed2",
		"--float64Slice", "13.3", "--float64Slice", "15.5",
		"--int64Slice", "13", "--int64Slice", "15",
		"--intSlice", "14", "--intSlice", "16",
	}

	app := cli.NewApp()
	app.Name = "multi_values"
	app.Flags = []cli.Flag{
		&cli.StringSliceFlag{Name: "stringSlice"},
		&cli.Float64SliceFlag{Name: "float64Slice"},
		&cli.Int64SliceFlag{Name: "int64Slice"},
		&cli.IntSliceFlag{Name: "intSlice"},
	}
	app.Action = func(ctx *cli.Context) error {
		for i, v := range ctx.FlagNames() {
			fmt.Printf("%d-%s %#v\n", i, v, ctx.Value(v))
		}

		fmt.Println()
		strs := ctx.StringSlice("stringSlice")
		for _, str := range strs {
			fmt.Println(str)
		}

		float64s := ctx.Float64Slice("float64Slice")
		for _, f := range float64s {
			fmt.Println(f)
		}

		int64s := ctx.Int64Slice("int64Slice")
		for _, i := range int64s {
			fmt.Println(i)
		}

		ints := ctx.IntSlice("intSlice")
		for _, i := range ints {
			fmt.Println(i)
		}

		err := ctx.Err()
		fmt.Println("error:", err)
		return err
	}

	_ = app.Run(os.Args)
}
