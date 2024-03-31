package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// User ...
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" binding:"required"`
	Age      int    `form:"age" json:"age"`
}

func main() {
	r := gin.New()

	// curl http://127.0.0.1:8080/
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!\n")
	})

	// curl http://127.0.0.1:8080/ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})

	// curl http://127.0.0.1:8080/user/alice
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s\n", name)
	})

	// curl http://127.0.0.1:8080/welcome
	// curl http://127.0.0.1:8080/welcome/   not match
	// curl -s 'http://127.0.0.1:8080/welcome?firstname=alice&lastname=alice'
	// curl -s 'http://127.0.0.1:8080/welcome?firstname=alice'
	// curl -s 'http://127.0.0.1:8080/welcome?lastname=alice'
	// curl -s 'http://127.0.0.1:8080/welcome?firstname=中国'
	// curl -s 'http://127.0.0.1:8080/welcome?lastname=中国'
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s\n", firstname, lastname)
	})

	// curl -X POST http://127.0.0.1:8080/form_post -H "Content-Type:application/x-www-form-urlencoded" -d 'message=alice&nick=alice'
	// curl -X POST http://127.0.0.1:8080/form_post -H "Content-Type:application/x-www-form-urlencoded" -d 'message=alice&nick=alice' | python -m json.tool
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick":    nick,
		})
	})

	// curl -s -X PUT http://127.0.0.1:8080/post?id=alice\&page=1 -d 'name=alice&message=alice'
	r.PUT("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
		})
	})

	// curl -X POST http://127.0.0.1:8080/upload -F 'upload=@/root/gopath/src/go-demos/gin-demo/alice.txt' -H "Content-Type:multipart/form-data"
	// curl -X POST http://127.0.0.1:8080/upload -d 'name=alice'
	r.POST("/upload", func(c *gin.Context) {
		// name := c.PostForm("name")
		// fmt.Println("name:", name)
		// 使用c.Request.FormFile解析客户端文件name属性
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		dir, _ := os.Getwd()
		filename := dir + "/" + header.Filename
		fmt.Println(filename)

		out, err := os.Create(filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "create file fail", err)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.String(http.StatusInternalServerError, "copy file fail", err)
			return
		}

		c.String(http.StatusCreated, "upload successful")
	})

	// curl -X POST http://127.0.0.1:8080/multi/upload -F 'upload=@/root/gopath/src/go-demos/gin-demo/alice.txt' -F 'upload=@/root/gopath/src/go-demos/gin-demo/alice1.txt' -H "Content-Type:multipart/form-data"
	r.POST("/multi/upload", func(c *gin.Context) {
		// 与单个文件上传类似，只不过使用了c.Request.MultipartForm得到文件句柄
		// 再获取文件数据，然后遍历读写
		err := c.Request.ParseMultipartForm(200000)
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request\n")
			return
		}

		formdata := c.Request.MultipartForm
		files := formdata.File["upload"]
		for i := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				c.String(http.StatusInternalServerError, "open file fail\n")
				return
			}

			out, err := os.Create(files[i].Filename)
			defer out.Close()
			if err != nil {
				c.String(http.StatusInternalServerError, "create file fail\n")
				return
			}

			_, err = io.Copy(out, file)
			if err != nil {
				c.String(http.StatusInternalServerError, "copy file fail\n")
				return
			}

			c.String(http.StatusCreated, "upload successful\n")
		}
	})

	// curl -X POST http://127.0.0.1:8080/login -d 'username=alice&passwd=passwd&age=24' -H "Content-Type:application/x-www-form-urlencoded"
	// curl -X POST http://127.0.0.1:8080/login -d '{"username":"alice", "passwd":"passwd", "age":24}' -H "Content-Type:application/json"
	r.POST("/login", func(c *gin.Context) {
		var user User
		var err error
		contentType := c.Request.Header.Get("Content-Type")

		switch contentType {
		case "application/json":
			err = c.BindJSON(&user)
		case "application/x-www-form-urlencoded":
			// err = c.BindWith(&user, binding.Form)
			err = c.ShouldBindWith(&user, binding.Form)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":   user.Username,
			"passwd": user.Passwd,
			"age":    user.Age,
		})
	})

	// 自动推断content-type是form还是json
	// curl -X POST http://127.0.0.1:8080/loginBind -d 'username=alice&passwd=passwd&age=24'
	// curl -X POST http://127.0.0.1:8080/loginBind -d '{"username":"alice", "passwd":"passwd", "age":24}' -H "Content-Type:application/json"
	r.POST("/loginBind", func(c *gin.Context) {
		var user User

		err := c.Bind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"passwd":   user.Passwd,
			"age":      user.Age,
		})
	})

	// curl http://127.0.0.1:8080/render?content_type=xml
	// curl http://127.0.0.1:8080/render?content_type=json
	r.GET("/render", func(c *gin.Context) {
		contentType := c.DefaultQuery("content_type", "json")
		if contentType == "json" {
			c.JSON(http.StatusOK, gin.H{
				"user":   "rsj217-json",
				"passwd": "123",
			})
		} else if contentType == "xml" {
			c.XML(http.StatusOK, gin.H{
				"user":   "rsj217-xml",
				"passwd": "123",
			})
		}
	})

	// 重定向
	// curl http://127.0.0.1:8080/redict/baidu
	r.GET("/redict/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://baidu.com")
	})

	// 全局中间件 注意在router.Use(MiddleWare())代码以上的路由函数，将不会有被中间件装饰的效果
	// curl http://127.0.0.1:8080/middleware
	// 使用花括号包含被装饰的路由函数只是一个代码规范，即使没有被包含在内的路由函数，
	// 只要使用router进行路由，都等于被装饰了。想要区分权限范围，可以使用组返回的对象注册中间件
	r.Use(MiddleWare())
	{
		r.GET("/middleware", func(c *gin.Context) {
			// 如果没有注册就使用MustGet方法读取c的值将会抛错，可以使用Get方法取而代之
			request := c.MustGet("request").(string)
			req, _ := c.Get("request")
			c.JSON(http.StatusOK, gin.H{
				"middile_request": request,
				"request":         req,
			})
		})
	}

	// 单个路由中间件
	r.GET("/before", MiddleWare(), func(c *gin.Context) {
		request := c.MustGet("request").(string)
		c.JSON(http.StatusOK, gin.H{
			"middile_request": request,
		})
	})

	// 鉴权中间件
	r.GET("/auth/signin", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    "123",
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.String(http.StatusOK, "Login successful")
	})

	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})

	// gin里可以借助协程实现异步任务。因为涉及异步过程，请求的上下文需要copy
	// 到异步的上下文，并且这个上下文是只读的
	// curl http://127.0.0.1:8080/sync
	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)
	})

	// curl http://127.0.0.1:8080/async
	r.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(2 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
	})

	r.Run("127.0.0.1:8080")
}

// MiddleWare 全局中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "client_request")
		c.Next()
		fmt.Println("after middleware")
	}
}

// AuthMiddleWare 鉴权中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}
