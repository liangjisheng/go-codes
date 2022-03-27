package lib1

import (
	"bou.ke/monkey"
	"encoding/json"
	"fmt"
	jjson "github.com/json-iterator/go"
	"github.com/smartystreets/goconvey/convey"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
)

type Dto struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Like string `json:"like"`
}

func TestPatchFunc(t *testing.T) {
	convey.Convey("case1", t, func() {
		monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
			s := make([]interface{}, len(a))
			for i, v := range a {
				s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
			}
			return fmt.Fprintln(os.Stdout, s...)
		})
		fmt.Println("what the hell?")
	})
}

func TestPatchMethod(t *testing.T) {
	convey.Convey("case1", t, func() {
		var d *net.Dialer
		monkey.PatchInstanceMethod(reflect.TypeOf(d), "Dial", func(_ *net.Dialer, _, _ string) (net.Conn, error) {
			return nil, fmt.Errorf("no dialing allowed")
		})
		_, err := http.Get("http://baidu.com")
		fmt.Println(err) // Get http://baidu.com: no dialing allowed
	})
}

func TestJsonPatch(t *testing.T) {
	monkey.Patch(json.Marshal, func(v interface{}) ([]byte, error) {
		fmt.Println("use jsoniter")
		return jjson.Marshal(v)
	})

	monkey.Patch(json.Unmarshal, func(data []byte, v interface{}) error {
		fmt.Println("use jsoniter")
		return jjson.Unmarshal(data, v)
	})
	dd := &Dto{
		Name: "xiaorui",
		Addr: "rfyiamcool@163.com",
		Like: "Python & Golang",
	}

	resDto := &Dto{}

	v, err := json.Marshal(dd)
	fmt.Println(string(v), err)

	errDe := json.Unmarshal(v, resDto)
	fmt.Println(resDto, errDe)

	fmt.Println("test end")
}

func TestPatchGuard(t *testing.T) {
	var guard *monkey.PatchGuard
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient), "Get", func(c *http.Client, url string) (*http.Response, error) {
		guard.Unpatch()
		defer guard.Restore()

		if !strings.HasPrefix(url, "https://") {
			return nil, fmt.Errorf("only https requests allowed")
		}

		return c.Get(url)
	})

	_, err := http.Get("http://baidu.com")
	fmt.Println(err) // only https requests allowed
	resp, err := http.Get("https://baidu.com")
	fmt.Println(resp.Status, err) // 200 OK <nil>
}
