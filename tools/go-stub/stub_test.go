package stubdemo

import (
	"fmt"
	"testing"
	"time"

	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

// 使用GoConvey测试框架和GoStub测试框架编写的测试用例如下
func TestFuncDemo(t *testing.T) {
	Convey("TestFuncDemo", t, func() {

		Convey("for success", func() {
			stubs := gostub.Stub(&counter, 200)
			defer stubs.Reset()

			stubs.Stub(&timeNow, func() time.Time {
				return time.Date(2020, 5, 19, 16, 4, 30, 0, time.UTC)
			})

			stubs.StubFunc(&Cleanup)

			fmt.Println(counter)
			fmt.Println(timeNow().Day())
			Cleanup("hello go")
		})

	})
}
