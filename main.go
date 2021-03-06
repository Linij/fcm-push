package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	controllers2 "go-push/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	mvc.New(app.Party("/register")).Handle(new(controllers2.RegisterController))
	mvc.New(app.Party("/fcm")).Handle(new(controllers2.FcmController))

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("0.0.0.0:8089"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
