package main

import (
	"fmt"
	"golang.org/x/text/currency"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

func init() {
	message.SetString(language.Chinese, "%s went to %s.", "%s去了%s。")
	message.SetString(language.AmericanEnglish, "%s went to %s.", "%s is in %s.")

	message.SetString(language.Chinese, "%s has been stolen.", "%s被偷走了。")
	message.SetString(language.AmericanEnglish, "%s has been stolen.", "%s has been stolen.")

	message.SetString(language.Chinese, "HOW_ARE_U", "%s 你好吗?")
	message.SetString(language.AmericanEnglish, "HOW_ARE_U", "%s How are you?")

	//根据参数处理不同的返回语句
	message.Set(language.English, "APP_COUNT",
		plural.Selectf(1, "%d",
			"=1", "I have an apple",
			"=2", "I have two apples",
			"other", "I have %[1]d apples",
		))

	// 以上代码和以下代码都是硬编码方式
	for _, e := range MSAArray {
		tag := language.MustParse(e.tag)
		switch msg := e.msg.(type) {
		case string:
			message.SetString(tag, e.key, msg)
		case catalog.Message:
			message.Set(tag, e.key, msg)
		case []catalog.Message:
			message.Set(tag, e.key, msg...)
		}
	}
}

type langMsg struct {
	tag, key string
	msg      interface{}
}

//手工 硬编码方式
var MSAArray = [...]langMsg{
	{"en", "HELLO_WORLD", "%s Hello World"},
	{"zh", "HELLO_WORLD", "%s 你好世界"},
	{"en", "TASK_REM", plural.Selectf(1, "%d",
		"=1", "One task remaining!",
		"=2", "Two tasks remaining!",
		"other", "[1]d tasks remaining!",
	)},
	{"zh", "TASK_REM", plural.Selectf(1, "%d",
		"=1", "剩余一项任务！",
		"=2", "剩余两项任务！",
		"other", "剩余 [1]d 项任务！",
	)},
}

func hardcode() {
	// 中文版
	p := message.NewPrinter(language.Chinese)
	p.Printf("%s went to %s.", "彼得", "英格兰")
	fmt.Println()
	p.Printf("%s has been stolen.", "宝石")
	fmt.Println()
	p.Printf("HOW_ARE_U", "竹子")
	fmt.Println()

	// 英文版本
	p = message.NewPrinter(language.AmericanEnglish)
	p.Printf("%s went to %s.", "Peter", "England")
	fmt.Println()
	p.Printf("%s has been stolen.", "The Gem")
	fmt.Println()
	p.Printf("HOW_ARE_U", "bamboo")
	fmt.Println()

	fmt.Println("placehold中的条件判断-------------------")
	// 条件判断
	p.Printf("APP_COUNT", 1)
	fmt.Println()
	p.Printf("APP_COUNT", 2)
	p.Println()
	p.Printf("APP_COUNT", 3)
	p.Println()

	fmt.Println("货币单位-------------------")
	// 货币单位
	p.Printf("%d", currency.Symbol(currency.USD.Amount(0.1)))//符号  美元货币格式化
	fmt.Println()
	p.Printf("%d", currency.NarrowSymbol(currency.CNY.Amount(1.6))) // 窄符号
	fmt.Println()
	p.Printf("%d", currency.ISO.Kind(currency.Cash)(currency.EUR.Amount(12.255)))//国际符号代码 欧元格式化
	fmt.Println()

	// 调用硬编码中的消息 国际化
	p = message.NewPrinter(language.English)
	p.Printf("HELLO_WORLD","bamboo")
	p.Println()
	p.Printf("TASK_REM", 2)
	p.Println()

	fmt.Println("国家语言简写格式-------------------")
	// 语言类型构建
	zh, _ := language.ParseBase("zh") // 语言
	CN, _ := language.ParseRegion("CN") // 地区
	zhLngTag, _ := language.Compose(zh, CN)
	fmt.Println(zhLngTag) // 打印 zh-CN
	fmt.Println(language.Chinese)// 打印中文缩写
	fmt.Println(language.SimplifiedChinese)// 打印中文缩写
	fmt.Println(language.TraditionalChinese)// 打印中文缩写
	fmt.Println(language.AmericanEnglish)// 打印英文缩写
}
