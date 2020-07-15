package calc

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	if ans := Mul(2, 3); ans != 6 {
		t.Errorf("2 * 3 expected be 6, but %d got", ans)
	}

	if ans := Mul(-2, 3); ans != -6 {
		t.Errorf("-2 * 3 expected be -6, but %d got", ans)
	}
}

// 子测试是 Go 语言内置支持的，可以在某个测试用例中，根据测试场景使用 t.Run创建不同的子测试用例
func TestMul1(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(-2, 3) != -6 {
			t.Fatal("fail")
		}
	})
}

// 对于多个子测试的场景，更推荐如下的写法(table-driven tests)
func TestMul2(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", -2, 3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}

// 对一些重复的逻辑，抽取出来作为公共的帮助函数(helpers)
// 可以增加测试代码的可读性和可维护性 借助帮助函数，可以让
// 测试用例的主逻辑看起来更清晰

type calcCase struct{ A, B, Expected int }

func createMulTestCase(t *testing.T, c *calcCase) {
	// Go 语言在 1.9 版本中引入了 t.Helper()，用于标注该函数是帮助函数
	// 报错时将输出帮助函数调用者的信息，而不是帮助函数的内部信息
	t.Helper()

	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got",
			c.A, c.B, c.Expected, ans)
	}
}

func TestMul3(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{-2, 3, -6})
	createMulTestCase(t, &calcCase{2, 0, 1})
}

// 如果在同一个测试文件中，每一个测试用例运行前后的逻辑是相同的
// 一般会写在 setup 和 teardown 函数中。例如执行前需要实例化待
// 测试的对象，如果这个对象比较复杂，很适合将这一部分逻辑提取出来
// 执行后可能会做一些资源回收类的工作 例如关闭网络连接 释放文件等

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// 测试某个 API 接口的 handler 能够正常工作

func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}

// 针对 http 开发的场景，使用标准库 net/http/httptest 进行测试更为高效
// 使用 httptest 模拟请求对象(req)和响应对象(w)，达到了相同的目的
func TestConn1(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)

	bytes, _ := ioutil.ReadAll(w.Result().Body)
	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello")
	}
}

// 使用 RunParallel 测试并发性能
func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
