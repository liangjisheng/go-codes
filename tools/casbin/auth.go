package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"strings"
)

var enforcer *casbin.Enforcer

const (
	UserRoleAdmin   = "admin"
	UserRoleVisitor = "visitor"
)

func InitCasbin(modelsFile string) {
	dsn := fmt.Sprintf("root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=UTC")
	adapter, err := gormadapter.NewAdapter("mysql", dsn, true)
	if err != nil {
		panic(err)
	}

	enforcer, err = casbin.NewEnforcer(modelsFile, adapter)
	if err != nil {
		panic(err)
	}

	//add default policy
	_, _ = enforcer.AddPolicy(UserRoleAdmin, "/api/*", "GET")
	_, _ = enforcer.AddPolicy(UserRoleAdmin, "/api/*", "POST")
	_, _ = enforcer.AddPolicy(UserRoleVisitor, "/api/*", "GET")
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}

func PermissionMiddleWare(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if ctx.FullPath() == "/admin/login" {
			ctx.Next()
			return
		}

		// 获取前端传回的token(传递方式不同，获取的位置也不同，根据实际情况选择)
		uuid := ctx.Request.Header.Get("zk-uuid")
		if uuid == "" {
			ctx.Abort()
			return
		}

		// get user info by uuid
		user := GetUserByUUID(uuid)
		sub := user.Role
		obj := ctx.FullPath()
		act := strings.ToUpper(ctx.Request.Method)
		//check the permission
		ok, _ := enforcer.Enforce(sub, obj, act)
		if !ok {
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

type User struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
	Role string `json:"role"`
}

func GetUserByUUID(uuid string) *User {
	//	actual get from db
	return &User{
		Name: "name",
		UUID: uuid,
		Role: UserRoleVisitor,
	}
}
