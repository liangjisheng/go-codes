package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io/ioutil"
)

func jsonfile() {
	fmt.Println("从配置文件中读取配置json解析并加载到系统中-------------------")
	InitConfig("messages_zh.json") // 加载配置信息json
	InitConfig("messages_en.json") // 加载配置信息json
	p := message.NewPrinter(language.SimplifiedChinese) // 设置语言类型
	p.Printf("HELLO_1", "Peter")
	fmt.Println()

	p.Printf("VISITOR", "Peter","用户管理系统接口")
	fmt.Println()

	p = message.NewPrinter(language.AmericanEnglish) // 设置语言类型
	p.Printf("VISITOR", "Peter","UER MANAGE SYSTEM API")
	fmt.Println()

	msg := Message{"en", "HELLO_WORLD", "%s Hello World"}
	p.Printf("CONFIG", msg) // 传一个对象值进去
	fmt.Println()
}

// 后面的jsoN字符串隐射在生成json字符串时需要,默认
type Message struct {
	Id string `json:"id"`
	Message string `json:"message,omitempty"`
	Translation   string `json:"translation,omitempty"`
}

type I18n struct {
	Language string `json:"language"`
	Messages []Message `json:"messages"`
}

// ioutil读写文件，依赖 io/ioutil 主要侧重文件和临时文件的读取和写入，对文件夹的操作较少
func ReadI18nJson(file string) string {
	b, err := ioutil.ReadFile(file)
	Check(err)
	str := string(b)
	return str
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func InitConfig(jsonPath string) {
	var i18n I18n
	str := ReadI18nJson(jsonPath)
	json.Unmarshal([]byte(str), &i18n)
	fmt.Println(i18n.Language)

	msaArray := i18n.Messages
	tag := language.MustParse(i18n.Language)
	// 以上代码和以下代码都是硬编码方式
	for _, e := range msaArray {
		fmt.Println(e.Id+"\t"+e.Translation)
		message.SetString(tag, e.Id, e.Translation)
	}
}
