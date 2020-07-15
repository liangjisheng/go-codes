package counter

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/host"
)

// Configurator ...
func Configurator(app *iris.Application) {
	counterValue := 0

	go func() {
		ticker := time.NewTicker(time.Second)

		for range ticker.C {
			counterValue++
		}

		app.ConfigureHost(func(h *host.Supervisor) {
			h.RegisterOnShutdown(func() {
				ticker.Stop()
			})
		})
	}()

	app.Get("/counter", func(ctx iris.Context) {
		ctx.Writef("Counter value = %d", counterValue)
	})
}
