package apis

import (
	"math"
	"net/http"
	"strconv"

	"go-demos/gin-demo/go-sample/models"

	"github.com/gin-gonic/gin"
)

// UserIndex ...
// curl -s -X GET http://127.0.0.1:8080/user/index?page=1\&size=5
func UserIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	data := models.GetUsers(page, size)

	meta := make(map[string]interface{})
	total, _ := models.GetUserTotal()
	meta["total"] = total
	meta["current_page"] = page
	meta["per_page"] = size
	meta["last_page"] = math.Ceil(float64(total/size)) + 1

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
		"meta": meta,
	})
}
