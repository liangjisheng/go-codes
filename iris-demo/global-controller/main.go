package main

import (
	"fmt"
	"sync/atomic"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&globalVisitorsController{visits: 0})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

type globalVisitorsController struct {
	//当访问单一控制器的时候，由开发人员负责并发的安全访问，
	// 因为所有的客户端共享一个相同的控制器实例。
	// 请注意任何控制器的方法都是每个客户端的，
	// 但是如果此结构没有任何动态结构字段依赖于 Iris ，该结构的字段可以在多个客户端共享。
	// Context 和 ALL 字段的值不为0，在这种情况下我们使用 unit64 ，
	// 它的值不是0（即使我们没有手动设置它），而是 &{0} 。
	// 以上所有都声明了只有一个 Singleton ，注意你不必去写代码去实现它，Iris 已经给你做好了。
	// 请看 `Get` 方法。
	visits uint64
}

func (c *globalVisitorsController) Get() string {
	count := atomic.AddUint64(&c.visits, 1)
	return fmt.Sprintf("Total visitors: %d", count)
}
