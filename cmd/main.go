package main

import (
	"github.com/carseven/go-fullstack-template/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	helloHandler := handler.HelloHandler{}
	app.GET("/hello", helloHandler.HandleHello)
	app.GET("/hello/:name", helloHandler.HandleHello)

	app.Start(":3000")

}
