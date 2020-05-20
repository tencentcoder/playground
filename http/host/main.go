package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/", func(ctx context.Context) {
		_, _ = ctx.Write([]byte("Hello World"))
	})
	if err := app.Run(
		iris.Addr("0.0.0.0:80"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	); err != nil {
		println("Start Error")
	}
}
