package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// UserInfo ...
type UserInfo struct {
	ID       int    `json:"id" gorm:"auto-increment"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

// ResponseData ...
type ResponseData struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var db *gorm.DB
var err error

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:ljs199711@tcp(localhost:3306)/study?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open db error", err)
	} else {
		fmt.Println("open db success")
	}
}

func main() {
	if db != nil {
		defer db.Close()
	}
	db.AutoMigrate(&UserInfo{}) // create table

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/login", login)
		v1.POST("/user/register", register)
		v1.GET("/user/get_info", getUserInfo)
	}
	router.Run()
}

// post form
func login(c *gin.Context) {
	tel := c.PostForm("tel")
	psd := c.PostForm("password")
	if len(tel) == 0 || len(psd) == 0 {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "账号或密码不能为空",
			Data:    "",
		}
		c.JSON(200, response)
	} else {
		var user UserInfo
		db.Where("tel=?", tel).First(&user)
		if user.Tel == "" {
			response := ResponseData{
				Code:    50001,
				Status:  "error",
				Message: "用户不存在",
				Data:    "",
			}
			c.JSON(200, response)
		} else {
			if user.Password == psd {
				response := ResponseData{
					Code:    200,
					Status:  "success",
					Message: "登录成功",
					Data:    "",
				}
				c.JSON(200, response)
			} else {
				response := ResponseData{
					Code:    50001,
					Status:  "error",
					Message: "密码错误",
					Data:    "",
				}
				c.JSON(200, response)
			}
		}
	}
}

// post form
func register(c *gin.Context) {
	tel := c.PostForm("tel")
	psd := c.PostForm("password")
	if len(tel) == 0 || len(psd) == 0 {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "账号或密码不能为空",
			Data:    "",
		}
		c.JSON(200, response)
		return
	}

	var user UserInfo
	db.Where("tel=?", tel).First(&user)
	if user.Tel == tel {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "手机号已注册",
			Data:    "",
		}
		c.JSON(200, response)
	} else {
		newUser := UserInfo{Tel: tel, Password: psd}
		db.Create(&newUser)
		response := ResponseData{
			Code:    200,
			Status:  "success",
			Message: "注册成功",
			Data:    "",
		}
		c.JSON(200, response)
	}
}

// get url
func getUserInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "参数错误",
			Data:    "",
		}
		c.JSON(200, response)
		return
	}

	var user UserInfo
	db.First(&user, id)
	if user.ID > 0 {
		response := ResponseData{
			Code:    200,
			Status:  "success",
			Message: "",
			Data:    user,
		}
		c.JSON(200, response)
	} else {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "用户不存在",
			Data:    "",
		}
		c.JSON(200, response)
	}
}

func checkErr() {
	if err != nil {
		fmt.Println(err)
		err = nil
	}
}
