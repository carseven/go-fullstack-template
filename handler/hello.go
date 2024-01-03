package handler

import (
	"github.com/carseven/go-fullstack-template/view/cases/hello"
	layout "github.com/carseven/go-fullstack-template/view/layouts"
	"github.com/labstack/echo/v4"
)

type HelloHandler struct {
	Environment string
}

func (h HelloHandler) HandleHello(c echo.Context) error {
	name := c.Param("name")
	userLanguage := "es" // TODO: Get Accept-Language from request header
	return render(c, layout.Base(hello.View(name, userLanguage), h.Environment))
}
