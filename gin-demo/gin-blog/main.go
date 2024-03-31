package main

import (
	"fmt"
	"go-demos/gin-demo/gin-blog/pkg/setting"
	"go-demos/gin-demo/gin-blog/routers"
	"log"
	"net/http"
)

func main() {
	// 没有热更新
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening %s", s.Addr)
	s.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and
	// download github.com/fvbock/endless
	// endless 热更新是采取创建子进程后，将原进程退出的方式，这点不符合守护进程的要求
	// endless.DefaultReadTimeOut = setting.ReadTimeout
	// endless.DefaultWriteTimeOut = setting.WriteTimeout
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	// server := endless.NewServer(endPoint, routers.InitRouter())
	// server.BeforeBegin = func(add string) {
	// 	log.Printf("Actual pid is %d\n", syscall.Getpid())
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Printf("Server err: %v\n", err)
	// }
}
