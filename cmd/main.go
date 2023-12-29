package main

import (
	"io"
	"net/http"

	"github.com/carseven/go-fullstack-template/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	helloHandler := handler.HelloHandler{}
	app.GET("/hello", helloHandler.HandleHello)
	app.GET("/hello/:name", helloHandler.HandleHello)

	// TODO: Refactor to proper handlers
	app.GET("ip", func(c echo.Context) error {
		resp, err := http.Get("http://ipecho.net/plain")
		if err != nil {
			c.NoContent(http.StatusInternalServerError)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.NoContent(http.StatusInternalServerError)
		}

		return c.HTML(http.StatusOK, string(body))
	})

	app.Use(middleware.GzipWithConfig(
		middleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				return false
			},
			MinLength: 2000,
		},
	))
	// TODO: Add generic form to access all assets (Aware to avoid leaks!!!)
	// TODO: Add cache control policy
	app.File("/assets/htmx", "assets/js/htmx/htmx@1.9.10.min.js")

	app.Start(":3000")

}
