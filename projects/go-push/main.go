package main

import (
	"gopush/gateway"
	"gopush/logic"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go gateway.InitWsServer()
	go logic.InitHttpServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
