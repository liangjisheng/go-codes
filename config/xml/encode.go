package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

// 为了少敲几个字符，声明了attrmap类型和start函数
type attrmap map[string]string // 属性的键值对容器

// start()用来构建开始节点
func start(tag string, attrs attrmap) xml.StartElement {
	var a []xml.Attr
	for k, v := range attrs {
		a = append(a, xml.Attr{xml.Name{"", k}, v})
	}
	return xml.StartElement{xml.Name{"", tag}, a}
}

func encode() {
	buffer := new(bytes.Buffer)
	enc := xml.NewEncoder(buffer)
	enc.Indent("", "    ") // 设置缩进，4个空格

	// 开始生成xml
	startExtension := start("extension", attrmap{"name": "rtp_multicast_page"})
	enc.EncodeToken(startExtension)
	startCondition := start("condition", attrmap{"field": "destination_number",
		"expression": "^pagegroup$|^7243$"})
	enc.EncodeToken(startCondition)
	startAction := start("action", attrmap{"application": "answer"})
	enc.EncodeToken(startAction)
	enc.EncodeToken(xml.CharData("raw text"))
	enc.EncodeToken(startAction.End())
	startAction = start("action", attrmap{"application": "esf_page_group"})
	enc.EncodeToken(startAction)
	enc.EncodeToken(startAction.End())
	enc.EncodeToken(startCondition.End())
	enc.EncodeToken(startExtension.End())

	// 写入xml
	enc.Flush()
	fmt.Println(buffer)
}
