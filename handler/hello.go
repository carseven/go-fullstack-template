package handler

import (
	"fmt"

	hello "github.com/carseven/go-fullstack-template/view"
	"github.com/labstack/echo/v4"
)

type HelloHandler struct{}

func (h HelloHandler) HandleHello(c echo.Context) error {
	name := c.Param("name")

	fmt.Println("Hello")

	return render(c, hello.View(name))
}
