package handler

import (
	layout "github.com/carseven/go-fullstack-template/view/layouts"
	"github.com/labstack/echo/v4"
)

type HelloHandler struct{}

func (h HelloHandler) HandleHello(c echo.Context) error {
	name := c.Param("name")

	return render(c, layout.Base(name))
}
