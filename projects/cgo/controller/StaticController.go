package controller

import (
	"go-demos/projects/cgo/cgo"
	"go-demos/projects/cgo/constant"
	"net/http"
)

// 静态资源
type StaticController struct {
	cgo.ApiController
}

func (p *StaticController) Router(router *cgo.RouterHandler) {
	router.Router(constant.STATIC_BASE_PATH, p.img)
}

var static = http.StripPrefix(constant.STATIC_BASE_PATH, http.FileServer(http.Dir(constant.BASE_IMAGE_ADDRESS)))

func (p *StaticController) img(w http.ResponseWriter, r *http.Request) {
	static.ServeHTTP(w, r)
}
