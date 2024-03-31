package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
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

	//上面是将基本策略写入 mysql 中，下面这句也可以将基本策略放到一个 csv 文件中,
	//就是基本策略存放的地方不一样而已
	//enforcer, err = casbin.NewEnforcer("rbac_models.conf", "rbac2.csv")

	//add default policy, 会在 mysql 中写入3条数据
	//表示一个 role 能对一个 obj(url资源) 进行什么样的操作(get or post)
	_, _ = enforcer.AddPolicy(UserRoleAdmin, "/api/admin/*", "GET")
	_, _ = enforcer.AddPolicy(UserRoleAdmin, "/api/admin/*", "POST")
	_, _ = enforcer.AddPolicy(UserRoleVisitor, "/api/article/*", "GET")
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
		ok, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !ok {
			ctx.AbortWithError(http.StatusForbidden, fmt.Errorf("403 forbid"))
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
