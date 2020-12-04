package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/hello", func(context iris.Context) {
		_, _ = context.WriteString("hello")
		_, _ = context.WriteString(" world")
	})

	_ = app.Listen(":880")
}
