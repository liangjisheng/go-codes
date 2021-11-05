package cgo

import (
	"fmt"
	"go-demos/projects/cgo/constant"
	"net/http"
	"strings"
)

var Router *RouterHandler = new(RouterHandler)

type RouterHandler struct {
}

var mux = make(map[string]func(http.ResponseWriter, *http.Request))

func (p *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 所有的请求进来都打印下
	fmt.Println(r.URL.Path)

	// 非静态资源请求
	if fun, ok := mux[r.URL.Path]; ok {
		fun(w, r)
		return
	}

	// 静态资源请求
	if strings.HasPrefix(r.URL.Path, constant.STATIC_BASE_PATH) {
		if fun, ok := mux[constant.STATIC_BASE_PATH]; ok {
			fun(w, r)
			return
		}
	}

	http.Error(w, "error URL:"+r.URL.String(), http.StatusBadRequest)
}

// 添加路由
func (p *RouterHandler) Router(relativePath string, handle func(http.ResponseWriter, *http.Request)) {
	mux[relativePath] = handle
}
