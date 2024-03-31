package stubdemo

import (
	"fmt"
	"testing"

	"github.com/prashantv/gostub"
)

var counter = 100

func TestStubGlobalVar(t *testing.T) {
	stubs := gostub.Stub(&counter, 200)
	// Reset方法将全局变量的值恢复为原值
	defer stubs.Reset()
	fmt.Println("counter:", counter)
}
