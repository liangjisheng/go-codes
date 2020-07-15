package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"

	"bou.ke/monkey"
)

func patchFunc() {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?")
}

func patchMethod() {
	var d *net.Dialer
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Dial", func(_ *net.Dialer, _, _ string) (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	_, err := http.Get("http://baidu.com")
	fmt.Println(err) // Get http://baidu.com: no dialing allowed
}
