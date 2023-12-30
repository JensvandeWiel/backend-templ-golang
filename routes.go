package main

import (
	"backend-templ-golang/app"
	"backend-templ-golang/controllers"
)

func setupRoutes(a *app.App) {
	helloWorldController := controllers.HelloWorldController{
		BaseController: app.BaseController{App: a},
	}

	a.Server.Static("/public", "public")

	a.Server.GET("/", helloWorldController.Index)
	a.Server.GET("/venus", helloWorldController.Venus)
}
