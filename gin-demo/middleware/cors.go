package middleware

//https://github.com/gin-contrib/cors

import (
	"github.com/gin-contrib/cors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		//接收客户端发送的origin
		c.Header("Access-Control-Allow-Origin", origin)
		//允许跨域设置可以返回其他子段，可以自定义字段
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		//服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		//允许浏览器（客户端）可以解析的头部
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		//允许客户端传递校验信息比如 cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		//设置缓存时间
		c.Header("Access-Control-Max-Age", "172800")

		// 放行所有OPTIONS方法
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
