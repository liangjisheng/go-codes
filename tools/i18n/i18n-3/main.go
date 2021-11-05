package main

import (
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"log"
)

func main() {
	i18n := gi18n.New(gi18n.Options{
		Language: "zh-CN",
	})


	i18n.SetPath("./testdata")

	helloCN := i18n.Translate(context.Background(), "hello")
	log.Println("helloCN:", helloCN)

	i18n.SetLanguage("en")
	helloEN := i18n.GetContent(context.Background(), "hello")
	log.Println("helloEN:", helloEN)

	i18n.SetLanguage("ja")
	helloJA := i18n.T(context.Background(), "{#hello}")
	log.Println("helloJA:", helloJA)

	i18n.SetLanguage("ru")
	helloRU := i18n.T(context.Background(), "{#hello}")
	log.Println("helloRU:", helloRU)
}
