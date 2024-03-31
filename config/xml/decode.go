package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// 低阶方法是以Token为单位操纵XML，Token有四种类型：StartElement，
// 用来表示XML开始节点；EndElement，用来表示XML结束节点；CharData，
// 即为XML的原始文本(raw text)；Comment，表示注释
// 该例用一个无限for循环，不断的获取Token，然后用Type Switch判断类型
// 根据不同的类型进行处理

func decode() {
	// 要解析的XML如下，为了提高可读性，用+号连接若干字符串，用以进行排版
	data :=
		`<extension name="rtp_multicast_page">` +
			`<condition field="destination_number" expression="^pagegroup$|^7243$">` +
			`<!-- comment -->` +
			`<action application="answer">raw text</action>` +
			`<action application="esf_page_group"/>` +
			`</condition>` +
			`</extension>`

	// 创建一个reader,以满足io.Reader接口
	reader := bytes.NewBuffer([]byte(data))
	dec := xml.NewDecoder(reader)

	// 开始遍历解码
	indent := ""  // 控制缩进
	sep := "    " // 每层缩进量为4个空格
	for {
		tok, err := dec.Token() // 返回下一个Token
		if err == io.EOF {      // 如果读到结尾，则退出循环
			break
		} else if err != nil { // 其他错误则退出程序
			os.Exit(1)
		}

		switch tok := tok.(type) { // Type switch
		case xml.StartElement: // 开始节点，打印名字和属性
			fmt.Print(indent)
			fmt.Printf("<%s ", tok.Name.Local)
			s := ""
			for _, v := range tok.Attr {
				fmt.Printf(`%s%s="%s"`, s, v.Name.Local, v.Value)
				s = " "
			}
			fmt.Println(">")
			indent += sep // 遇到开始节点，则增加缩进量
		case xml.EndElement: // 结束节点，打印名字
			indent = indent[:len(indent)-len(sep)] // 遇到结束节点，则减少缩进量
			fmt.Printf("%s</%s>\n", indent, tok.Name.Local)
		case xml.CharData: // 原始字符串，直接打印
			fmt.Printf("%s%s\n", indent, tok)
		case xml.Comment: // 注释，直接打印
			fmt.Printf("%s<!-- %s -->\n", indent, tok)
		}
	}
}
