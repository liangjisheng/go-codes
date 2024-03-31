package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// RemoteAddress ...
type RemoteAddress struct {
	TestAddr       []string `mapstructure:"test_addr"`
	ProductionAddr []string `mapstructure:"production_addr"`
}

// HTTPConfig ...
type HTTPConfig struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

// RPCConfig ...
type RPCConfig struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

// LogConfig ...
type LogConfig struct {
	Dir string `mapstructure:"dir"`
}

// Server ...
type Server struct {
	HTTP          HTTPConfig    `mapstructure:"http"`
	RPC           RPCConfig     `mapstructure:"rpc"`
	RemoteAddress RemoteAddress `mapstructure:"remote_address"`
	LogConfig     LogConfig     `mapstructure:"log"`
}

// Client ...
type Client struct {
	HTTP      HTTPConfig `mapstructure:"http"`
	RPC       RPCConfig  `mapstructure:"rpc"`
	LogConfig LogConfig  `mapstructure:"log"`
}

// Config ...
type Config struct {
	Server Server `mapstructure:"server"`
	Client Client `mapstructure:"client"`
}

func main() {
	c := &Config{}
	filename := "./config.yaml"

	viper.SetConfigType("yaml")
	viper.SetConfigFile(filename)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config is failed err:", err)
		os.Exit(-1)
	}

	err = viper.Unmarshal(c)
	if err != nil {
		fmt.Println("unmarshal config is failed, err:", err)
		os.Exit(-1)
	}

	fmt.Printf("%+v", *c)
}
