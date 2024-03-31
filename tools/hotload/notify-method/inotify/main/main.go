package main

import (
	"chainup/hotload/inotify/watcher"
	"fmt"
	"log"
	"syscall"
)

func main() {
	path := "./env.json"
	notify, err := watcher.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = notify.AddWatcher(path, syscall.IN_CLOSE_WRITE)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool, 1)
	go func() {
		for {
			select {
			case event := <-notify.Events:
				if event&syscall.IN_CLOSE_WRITE == syscall.IN_CLOSE_WRITE {
					fmt.Printf("file changed \n")
				}
			}
		}
	}()

	<-done
}
