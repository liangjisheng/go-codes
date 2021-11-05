package main

import (
	"github.com/syyongx/ii18n"
	"log"
)

func main() {
	config := map[string]ii18n.Config{
		"app": {
			SourceNewFunc: ii18n.NewJSONSource,
			OriginalLang: "en-US",
			BasePath: "./testdata",
			FileMap: map[string]string{
				"app": "app.json",
				"error": "error.json",
			},
		},
	}

	ii18n.NewI18N(config)

	helloEN := ii18n.T("app", "hello", nil, "en-US")
	log.Println("message:", helloEN)
	helloZH := ii18n.T("app", "hello", nil, "zh-CN")
	log.Println("message:", helloZH)
	nice := ii18n.T("app", "nice", nil, "zh-CN")
	log.Println("message:", nice)
	errorStr := ii18n.T("error", "error", nil, "zh-CN")
	log.Println("message:", errorStr)
	warn := ii18n.T("error", "warning", nil, "zh-CN")
	log.Println("message:", warn)
}
