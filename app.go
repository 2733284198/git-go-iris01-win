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

		//context.WriteString("hello user" + context.Params().Get("id"))
		fmt.Println(context.Params().Get("id"))

		context.JSON(iris.Map{
			"name": " hello  iris",
		})
	})

	app.PartyFunc("/cpanel", func(child iris.Party) {
		child.Get("/", func(ctx iris.Context) {
			ctx.WriteString("hello")
		})

		child.Get("/hi", func(ctx iris.Context) {
			ctx.WriteString("cpanel hi")
		})
	})

	app.PartyFunc("/test", func(child iris.Party) {
		app.Get("/{username:string}", profileByUsername)

		child.Get("/file", func(ctx iris.Context) {
			file := "./files/1.txt"
			ctx.SendFile(file, "2.txt")
		})
	})

	//app.Get("/profile/{username:string}", profileByUsername)
	//app.Get("/profile/{username:string}", profileByUsername)
	//app.Get("/test/file", sendfile)

	app.Run(iris.Addr(":8080"))
	//_ = app.Listen(":880")
}

func profileByUsername(ctx iris.Context) {
	//获取路由参数
	username := ctx.Params().Get("username")
	//向数据模板传值 当然也可以绑定其他值
	ctx.ViewData("Username", username)
	//渲染模板 ./web/views/profile.html

	//把获得的动态数据username 绑定在 ./web/views/profile.html 模板 语法{{}} {{ .Username }}

	ctx.View("profile.html")
}
