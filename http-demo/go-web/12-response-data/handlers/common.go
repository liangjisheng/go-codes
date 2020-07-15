package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Greeting ...
type Greeting struct {
	Content string `json:"greeting"`
}

// Home ...
func Home(w http.ResponseWriter, r *http.Request) {
	// Write 方法用于写入数据到 HTTP 响应实体, 如果调用 Write 方法时还不知道 Content-Type,
	// 会通过数据的前 512 个字节进行判断
	// io.WriteString(w, "Welcome to my blog site")

	// Content-Type 会自动调整成了 text/plain
	// w.Write([]byte("欢迎访问学院君个人网站👏"))

	// Content-Type 会自动调整成了 text/html
	// html := `<html>
	//     <head>
	//         <title>学院君个人网站</title>
	//     </head>
	//     <body>
	//         <h1>欢迎访问学院君个人网站👏</h1>
	//     </body>
	// </html>`
	// w.Write([]byte(html))

	// 返回 JSON 格式数据
	greeting := Greeting{
		"欢迎访问学院君个人网站👏",
	}
	message, _ := json.Marshal(greeting)
	// 要返回json这个格式的响应头, 需要设置响应头才能实现
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

// Error ...
func Error(w http.ResponseWriter, r *http.Request) {
	// WriteHeader 这个方法名有点误导, 其实它并不是用来设置响应头的, 该方法支持传入
	// 一个整型数据用来表示响应状态码, 如果不调用该方法的话, 默认响应状态码是 200 OK
	w.WriteHeader(401)
	fmt.Fprintln(w, "认证后才能访问该接口")
}

// Redirect ...
func Redirect(w http.ResponseWriter, r *http.Request) {
	// 对于重定向请求, 无需设置响应实体, 另外需要注意的是 w.Header().Set
	// 必须在 w.WriteHeader 之前调用, 因为一旦调用 w.WriteHeader 之后, 就不能对响应头进行设置了

	// 设置响应头 设置一个 301 重定向
	w.Header().Set("Location", "https://www.baidu.com")
	// 返回状态码
	w.WriteHeader(301)
}
