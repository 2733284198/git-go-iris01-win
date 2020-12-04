package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/hello/:id", func(context iris.Context) {

		//_, _ = context.WriteString("hello")
		//_, _ = context.WriteString(" world")

		context.WriteString("hello user" + context.Params().Get("id"))
		fmt.Println(context.Params().Get("id"))

		context.JSON(iris.Map{
			"name": " hello  iris",
		})
	})

	_ = app.Listen(":880")
}
