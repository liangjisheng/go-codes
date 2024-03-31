package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
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

// SetCookie ...
func SetCookie(w http.ResponseWriter, r *http.Request) {
	// 有两种方法来设置过期时间: 一种是直接设置 Expires 字段, 一种是设置 MaxAge 字段,
	// 前者表示到期的具体时间点, 后者表示 Cookie 的有效时长(单位是秒). 这并不是 Go 语言的设计,
	// 而是不同浏览器的混乱标准使然
	c1 := http.Cookie{
		Name:     "username",
		Value:    url.QueryEscape("学院君"),
		HttpOnly: true,
		// 如果用 Expires 字段来设置的话, 可以设置 Unix 时间戳的值为1
		// (对应的绝对时间是 1970-01-01 08:00:01 +0800 CST, 也就是一个过去的时间)
		Expires: time.Now().AddDate(0, 0, 1), // Cookie 有效期设置为1天
	}
	c2 := http.Cookie{
		Name:     "website",
		Value:    "https://xueyuanjun.com",
		HttpOnly: true,
		// 如果想要在 Cookie 过期之前提前删除 Cookie, 可以将 MaxAge 设置为小于 0 的值即可
		MaxAge: 1000, // Cookie 有效期设置为 1000s
	}
	w.Header().Add("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	// 当然也可以通过 http.SetCookie 方法写入 Cookie 到 HTTP 响应来实现, 这样做更便捷
	// http.SetCookie(w, &c1)
	// http.SetCookie(w, &c2)
	fmt.Fprintln(w, "通过 HTTP 响应头发送 Cookie 信息")
}

// GetCookie ...
func GetCookie(w http.ResponseWriter, r *http.Request) {
	// 要在服务端获取这些 Cookie 信息, 可以通过读取请求头的方式
	// cookie := r.Header.Get("Cookie")
	// 但是这种方式读取的 Cookie 字符串值还需要进行解析, 才能得到每个 Cookie 的值,
	// 为此可以通过更加便捷的专门用于读取每个 Cookie 的 r.Cookie 方法

	c1, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "名为 username 的 Cookie 不存在")
		return
	}
	username, _ := url.QueryUnescape(c1.Value)
	c2, err := r.Cookie("website")
	if err != nil {
		fmt.Fprintln(w, "名为 website 的 Cookie 不存在")
		return
	}
	website := c2.Value
	fmt.Fprintf(w, "从用户请求中读取的 Cookie: {username: %s, website: %s}\n", username, website)

	// 如果想要一次性获取所有 Cookie, 还可以通过 r.Cookies() 方法
	// 	cookies := r.Cookies()
	// c1 := cookies[0]  // username=%E5%AD%A6%E9%99%A2%E5%90%9B
	// c2 := cookies[1]  // website=https://xueyuanjun.com
}

// 使用 Cookie 设置一次性消息, 所谓一次性消息, 指的是页面重新加载后消息就不存在了,
// 也就是该消息只能被读取一次, 不管你用不用它都不复存在了. 我们可以结合上面的删除
// Cookie 功能来实现这个一次性消息功能

// SetWelcomeMessage ...
func SetWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	msg := "欢迎访问学院君网站👏"
	cookie := http.Cookie{
		Name:    "welcome_message",
		Value:   base64.URLEncoding.EncodeToString([]byte(msg)),
		Expires: time.Now().AddDate(0, 0, 1),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/get_welcome_message", 302)
}

// GetWelcomeMessage ...
func GetWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("welcome_message")
	if err != nil {
		fmt.Fprintln(w, "没有在 Cookie 中找到欢迎消息")
	} else {
		// 新增了读取成功之后, 删除这个 Cookie
		delCookie := http.Cookie{
			Name:   "welcome_message",
			MaxAge: -1,
		}
		http.SetCookie(w, &delCookie)
		msg, _ := base64.URLEncoding.DecodeString(cookie.Value)
		fmt.Fprintln(w, string(msg))
	}
}
