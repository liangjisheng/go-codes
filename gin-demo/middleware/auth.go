package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"strings"
)

func Auth(ctx *gin.Context) {
	if excludeAuth(ctx) {
		ctx.Next()
		return
	}

	bearer := ctx.Request.Header.Get("Authorization")
	token := strings.Split(bearer, "Bearer ")
	if len(token) != 2 {
		ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code": 400,
			"msg":  "should use Bearer Token",
			"data": nil,
		})
		ctx.Abort()
		return
	}

	claim, err := ParseToken(token[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code": 400,
			"msg":  fmt.Sprintf("auth parse failed, err=%v", err),
			"data": nil,
		})
		ctx.Abort()
		return
	}

	ctx.Set("address", claim.Address)
	ctx.Next()
}

func excludeAuth(ctx *gin.Context) bool {
	if ctx.Request.Method == "OPTIONS" {
		return true
	}

	s := ctx.FullPath()
	if s == "" ||
		s == "/" ||
		s == "/metrics" ||
		s == "/hello" {
		return true
	}

	return false
}
