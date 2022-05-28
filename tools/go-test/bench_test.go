package gotest

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	b.ResetTimer() // 重置性能测试计数
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello")
	}
}

// 使用 RunParallel 测试并发性能
func BenchmarkParallel(b *testing.B) {
	temp := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			temp.Execute(&buf, "World")
		}
	})
}
