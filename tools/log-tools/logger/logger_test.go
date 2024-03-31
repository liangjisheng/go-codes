package logger

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	WithTrace(context.Background()).Info("info")
	//for {
	//Info("info")
	//Debug("debug")
	//Warn("warn")
	//Error("error")
	//time.Sleep(200 * time.Millisecond)
	//}
}

func BenchmarkLog(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Info("info")
	}
}
