# lark

```shell
curl -X POST -H "Content-Type: application/json" -d '{"msg_type":"text","content":{"text":"request example"}}' "https://open.larksuite.com/open-apis/bot/v2/hook/xxxxxxxxxxxxxxxxx"
```

发送普通文本

```json
{
    "msg_type": "text",
    "content": {
        "text": "request example"
    }
}
```

发送富文本

```json
{
	"msg_type": "post",
	"content": {
		"post": {
			"zh_cn": {
				"title": "项目更新通知",
				"content": [
					[{
							"tag": "text",
							"text": "项目有更新: "
						},
						{
							"tag": "a",
							"text": "请查看",
							"href": "http://www.example.com/"
						},
						{
							"tag": "at",
							"user_id": "ou_18eac8********17ad4f02e8bbbb"
						}
					]
				]
			}
		}
	}
}
```
