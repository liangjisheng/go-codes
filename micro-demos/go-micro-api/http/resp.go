package http

import (
	"demo/utility/helper"

	"github.com/valyala/fasthttp"
)

// APIRet ...
func APIRet(ctx *fasthttp.RequestCtx, r Response) {
	b, _ := helper.JSONEncode(r)
	_, _ = ctx.Write(b)
}
