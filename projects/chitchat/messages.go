package main

import "github.com/nicksnyder/go-i18n/v2/i18n"

var messages = []i18n.Message{
	i18n.Message{
		ID:          "thread_not_found", // 消息文本的唯一标识
		Description: "Thread not exists in db",
		Other:       "Cannot read thread", // 对应的翻译字符串(默认是英文)
	},
}
