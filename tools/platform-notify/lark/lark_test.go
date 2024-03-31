package lark

import "testing"

func TestText(t *testing.T) {
	err := Text("test")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestRichText(t *testing.T) {
	content := [][]map[string]interface{}{
		{
			{
				"tag":  "text",
				"text": "第一行 :",
			},
			{
				"tag":  "a",
				"href": "https://www.google.com/",
				"text": "超链接",
			},
		},

		{
			{
				"tag":  "text",
				"text": "第二行:",
			},
			{
				"tag":  "text",
				"text": "文本测试",
			},
		},
	}

	err := RichText("rich text", content)
	if err != nil {
		t.Error(err)
		return
	}
}
