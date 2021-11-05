package main

import (
	"bookstore.example.com/internal/localizer"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"net/http"

	// 引入internal/translations包，确保init()函数被调用
	_ "bookstore.example.com/internal/translations"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	// 从URL路径提取语言区域，需根据你选择的路由实现方式修改这行代码.
	locale := r.URL.Query().Get(":locale")

	// Declare variable to hold the target language tag.
	var lang language.Tag

	// 使用 language.MustParse()为区域设置分配适当的语言标签。
	switch locale {
	case "en-gb":
		lang = language.MustParse("en-GB")
	case "de-de":
		lang = language.MustParse("de-DE")
	case "zh-cn":
		lang = language.MustParse("zh-CN")
	default:
		http.NotFound(w, r)
		return
	}

	// 使用对应语言初始化一个message.Printer实例
	p := message.NewPrinter(lang)
	// 将欢迎信息翻译成目标语言。
	p.Fprintf(w, "Welcome!\n")

	// 定义一个变量来保存书的数量。 在实际应用中数量需查询数据库的到
	var totalBookCount = 1
	//使用Fprintf() 函数在响应中添加书数量
	p.Fprintf(w, "%d books available\n", totalBookCount)
}

func handleHomeLocalizer(w http.ResponseWriter, r *http.Request) {
	// 基于URL中的区域设置ID初始化一个新的本地化器
	l, ok := localizer.Get(r.URL.Query().Get(":locale"))
	if !ok {
		http.NotFound(w, r)
		return
	}

	var totalBookCount = 1_252_794

	// 使用Translate()方法.
	fmt.Fprintln(w, l.Translate("Welcome!"))
	fmt.Fprintln(w, l.Translate("%d books available", totalBookCount))

	//增加 "Launching soon!"消息.
	fmt.Fprintln(w, l.Translate("Launching soon!"))
}