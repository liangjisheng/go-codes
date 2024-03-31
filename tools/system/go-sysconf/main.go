package main

import (
	"fmt"

	"github.com/tklauser/go-sysconf"
)

func main() {
	// get clock ticks, this will return the same as C.sysconf(C._SC_CLK_TCK)
	clktck, err := sysconf.Sysconf(sysconf.SC_CLK_TCK)
	if err == nil {
		fmt.Printf("SC_CLK_TCK: %v\n", clktck)
	}
}
