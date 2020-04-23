// file: main.go

package main

import (
	controllers2 "gcm-push/web/controllers"
	"github.com/kataras/iris/v12/_examples/mvc/overview/datasource"
	"github.com/kataras/iris/v12/_examples/mvc/overview/repositories"
	"github.com/kataras/iris/v12/_examples/mvc/overview/services"
	"github.com/kataras/iris/v12/_examples/mvc/overview/web/controllers"
	"github.com/kataras/iris/v12/_examples/mvc/overview/web/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	//app.Handle("GET", "/", func(ctx iris.Context) {
	//	// 导航到中间 $GOPATH/src/github.com/kataras/iris/context/context.go
	//	// 概述所有上下文的方法（有很多这样的方法，阅读它，你将学习iris如何工作
	//	ctx.HTML("Hello from " + ctx.Path()) // Hello from /
	//})
	//
	//app.RegisterView(iris.HTML("./web/views", ".html"))

	mvc.New(app.Party("/register")).Handle(new(controllers2.RegisterController))
	mvc.New(app.Party("/fcm")).Handle(new(controllers2.FcmController))

	// Serve our controllers.
	//mvc.New(app.Party("/hello")).Handle(new(controllers.HelloController))
	// You can also split the code you write to configure an mvc.Application
	// using the `mvc.Configure` method, as shown below.
	//mvc.Configure(app.Party("/movies"), movies)

	// http://localhost:8080/hello
	// http://localhost:8080/hello/iris
	// http://localhost:8080/movies
	// http://localhost:8080/movies/1
	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("192.168.2.97:8080"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

// note the mvc.Application, it's not iris.Application.
func movies(app *mvc.Application) {
	// Add the basic authentication(admin:password) middleware
	// for the /movies based requests.
	app.Router.Use(middleware.BasicAuth)

	// Create our movie repository with some (memory) data from the datasource.
	repo := repositories.NewMovieRepository(datasource.Movies)
	// Create our movie service, we will bind it to the movie app's dependencies.
	movieService := services.NewMovieService(repo)
	app.Register(movieService)

	// serve our movies controller.
	// Note that you can serve more than one controller
	// you can also create child mvc apps using the `movies.Party(relativePath)` or `movies.Clone(app.Party(...))`
	// if you want.
	app.Handle(new(controllers.MovieController))
}
