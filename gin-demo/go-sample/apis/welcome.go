package apis

import (
	"go-demos/gin-demo/go-sample/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome ...
// curl http://127.0.0.1:8080/welcome
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": pkg.GetUserSession(c),
	})
}
