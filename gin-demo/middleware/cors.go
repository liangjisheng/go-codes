package middleware

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//https://github.com/gin-contrib/cors
//https://juejin.cn/post/6844904100035821575#heading-67

//浏览器遵循同源政策(scheme(协议)、host(主机)和port(端口)都相同则为同源)
//当浏览器向目标 URI 发 Ajax 请求时，只要当前 URL 和目标 URL 不同源，则产生跨域，被称为跨域请求
//跨域请求的响应一般会被浏览器所拦截，注意，是被浏览器拦截，响应其实是成功到达客户端了

//在服务端处理完数据后，将响应返回，主进程检查到跨域，且没有cors(后面会详细说)响应头，
//将响应体全部丢掉，并不会发送给渲染进程。这就达到了拦截数据的目的。

//CORS 其实是 W3C 的一个标准，全称是跨域资源共享。它需要浏览器和服务器的共同支持
//服务器需要附加特定的响应头

//浏览器根据请求方法和请求头的特定字段，将请求做了一下分类, 分为简单请求和非简单请求

//简单请求
//请求方法为 GET、POST 或者 HEAD
//请求头的取值范围: Accept、Accept-Language、Content-Language、
//Content-Type(只限于三个值application/x-www-form-urlencoded、multipart/form-data、text/plain)

//除了简单请求剩下的就是非简单请求
//简单请求发出去之前，浏览器它会自动在请求头当中，添加一个 Origin 字段，用来说明请求来自哪个源。
//服务器拿到请求之后，在回应时对应地添加 Access-Control-Allow-Origin 字段，如果 Origin
//不在这个字段的范围中，那么浏览器就会将响应拦截

//执行非简单请求, 首先会发送预检请求
//预检请求内容
//OPTIONS / HTTP/1.1
//Origin: 当前地址
//Host: dst.com
//Access-Control-Request-Method: PUT
//Access-Control-Request-Headers: X-Custom-Header

//Access-Control-Request-Method, 列出 CORS 请求用到哪个HTTP方法
//Access-Control-Request-Headers，指定 CORS 请求将要加上什么请求头

//在预检请求的响应返回后，如果请求不满足响应头的条件, 真正的CORS请求也不会发出去了
//CORS 请求现在它和简单请求的情况是一样的。浏览器自动加上Origin字段，服务端响应头返回Access-Control-Allow-Origin

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//接收客户端发送的 origin, 不过 origin 一般直接写 *, 表示允许不同的源访问
		//scheme + host(domain) + port 组成一个源, 3 个当中有一个不同则表示不同的源
		//服务端将 origin 写入 response header, 浏览器拿到后判断请求源在这个响应头字段内
		//则不会拦截响应, 这个字段就是服务器用来决定浏览器是否拦截这个响应
		//origin := c.Request.Header.Get("Origin")
		//c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Origin", "*")

		//允许跨域请求发送的请求头字段
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		//服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH,HEAD")
		//允许浏览器（客户端）可以解析的头部, 不仅可以拿到基本的 6 个响应头字段
		//(包括Cache-Control、Content-Language、Content-Type、Expires、Last-Modified和Pragma)
		//还能拿到这个字段声明的响应头字段
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Type, Content-Language")
		//允许客户端传递校验信息比如 cookie
		//表示是否允许发送 Cookie，对于跨域请求，浏览器对这个字段默认值设为 false，而如果需要拿到浏览器的
		//Cookie，需要添加这个响应头并设为true, 并且在前端也需要设置withCredentials属性
		c.Header("Access-Control-Allow-Credentials", "true")
		//设置缓存时间, 2 days, 预检请求的有效期，在此期间，不用发出另外一条预检请求
		c.Header("Access-Control-Max-Age", "172800")

		// 放行所有 OPTIONS 和 HEAD 方法
		method := c.Request.Method
		if method == "OPTIONS" || method == "HEAD" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

//Cors1 使用 cors 默认配置的方法
func Cors1() gin.HandlerFunc {
	return cors.Default()
}
