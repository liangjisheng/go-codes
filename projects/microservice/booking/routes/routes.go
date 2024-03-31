package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-demos/projects/microservice/booking/controllers"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	MiddleWare mux.MiddlewareFunc // 中间件函数
}

var routes []Route

func init() {
	register("POST", "/booking", controllers.CreateBooking, nil)
	register("GET", "/booking", controllers.GetAllBooking, nil)
	register("GET", "/booking/{name}", controllers.GetBookByName, nil)
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).Path(route.Pattern)
		if route.MiddleWare != nil { // 使用中间件函数
			r.Handler(route.MiddleWare(route.Handler))
		} else { // 不使用中间件函数
			r.Handler(route.Handler)
		}
	}
	return router
}
