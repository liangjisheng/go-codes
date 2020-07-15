package pkg

import (
	"net/http"

	"go-demos/gin-demo/go-sample/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// KEY ...
const KEY = "AEN233"

// EnableCookieSession 使用 Cookie 保存 session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("SAMPLE", store)
}

// AuthSessionMiddle session 中间件
func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		if nil == sessionValue {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		// 设置简单的变量
		c.Set("userId", sessionValue.(uint))

		c.Next()
		return
	}
}

// SaveAuthSession 注册和登陆时都需要保存seesion信息
func SaveAuthSession(c *gin.Context, id uint) {
	session := sessions.Default(c)
	session.Set("userId", id)
	session.Save()
}

// ClearAuthSession 退出时清除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// HasSession ...
func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("userId"); sessionValue == nil {
		return false
	}
	return true
}

// GetSessionUserID ...
func GetSessionUserID(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return 0
	}
	return sessionValue.(uint)
}

// GetUserSession ...
func GetUserSession(c *gin.Context) map[string]interface{} {
	hasSession := HasSession(c)
	userName := ""
	if hasSession {
		userID := GetSessionUserID(c)
		userName = models.UserDetail(userID).Name
	}

	data := make(map[string]interface{})
	data["hasSession"] = hasSession
	data["userName"] = userName
	return data
}
