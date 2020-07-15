package config

import (
	"encoding/json"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

// ViperConfig ...
var ViperConfig Configuration

func init() {
	runtimeViper := viper.New()
	runtimeViper.AddConfigPath(".")
	runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("json")
	err := runtimeViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err.Error()))
	}
	runtimeViper.Unmarshal(&ViperConfig)

	// 本地化初始设置
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active.en.json")
	bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active." + ViperConfig.App.Language + ".json")
	ViperConfig.LocaleBundle = bundle

	// 监听配置文件变更 通过 runtimeViper.WatchConfig() 方法监听配置文件变更
	// (该监听会开启新的协程执行, 不影响和阻塞当前协程), 一旦配置文件有变更
	// 即可通过定义在 runtimeViper.OnConfigChange 中的匿名回调函数重新加载配置
	// 文件并将配置值映射到 ViperConfig 指针, 同时再次加载新的语言文件
	runtimeViper.WatchConfig()
	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
		runtimeViper.Unmarshal(&ViperConfig)
		ViperConfig.LocaleBundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active." + ViperConfig.App.Language + ".json")
	})
}
