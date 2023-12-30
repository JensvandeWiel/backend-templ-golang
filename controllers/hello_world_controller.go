package controllers

import (
	"backend-templ-golang/app"
	"backend-templ-golang/services"
	"github.com/labstack/echo/v4"
)

type HelloWorldController struct {
	app.BaseController
}

type HelloWorldPageData struct {
	Word string
}

func (controller *HelloWorldController) Index(c echo.Context) error {
	return c.Render(200, "hello_world", HelloWorldPageData{
		Word: "Mars",
	})
	//return c.String(200, services.NewHelloWorldService().HelloVenus())
}

func (controller *HelloWorldController) Venus(c echo.Context) error {
	return c.String(200, services.NewHelloWorldService().HelloVenus())
}
