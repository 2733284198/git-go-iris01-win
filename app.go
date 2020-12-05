package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	//"github.com/kataras/iris/v12/core/router"
)

func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/all", before, mainHandler, after)

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
			//ctx.WriteString("cpanel hi")
			ctx.HTML("<h1>Welcome</h1>")
		})

		child.Get("/param", func(ctx iris.Context) {
			//ctx.WriteString("cpanel hi")

			id, _ := ctx.Params().GetInt("id")

			//id := ctx.Params().Get("id")
			name := ctx.Params().Get("name")

			//id := ctx.Params().GetTrim("id")
			//id := ctx.Params().GetTrim("id")

			//fmt.Println(id)
			//fmt.Printf("%T, %d", id, id)
			//fmt.Printf("%T, %s", id, id)
			fmt.Printf("%T, %d, %T, %s", id, id, name, name)

			//strparam := fmt.Sprintf("%T, %d", id, id)
			//strparam := fmt.Sprintf("%T, %s", id, id)
			strparam := fmt.Sprintf("%T, %d, %T , %s", id, id, name, name)

			_, _ = ctx.WriteString(strparam)
		})
	})

	app.PartyFunc("/test", func(child iris.Party) {
		app.Get("/{username:string}", profileByUsername)

		child.Get("/file", func(ctx iris.Context) {
			file := "./files/1.txt"
			ctx.SendFile(file, "2.txt")
		})

		app.Get("/html", func(ctx iris.Context) {
			//ctx.HTML(" <h1>hi, I just exist in order to see if the server is closed</h1>")
			ctx.WriteString(" <h1>hi, I just exist in order to see if the server is closed</h1>")
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

// handle ,before , handle, after

func before(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() //继续执行下一个handler，在本例中是mainHandler。
}

func after(ctx iris.Context) {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")
	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")
	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)
	ctx.Next() // 继续下一个handler 这里是after
}
