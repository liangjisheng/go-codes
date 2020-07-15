package adminuser

import (
	"demo/http"

	adminuser "demo/app/api.gateway/service/admin.user"
	"github.com/valyala/fasthttp"
)

// UserLogin ...
func UserLogin(ctx *fasthttp.RequestCtx) {
	var (
		username = ctx.QueryArgs().Peek("username")
		password = ctx.QueryArgs().Peek("password")
	)

	err := adminuser.LoginByUserName(ctx, username, password)
	if err != nil {
		http.ApiRet(ctx, err)
		return
	}

	http.ApiRet(ctx, http.NewSuccess(nil))
}
