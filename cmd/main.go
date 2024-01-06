package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/carseven/go-fullstack-template/handler"
	"github.com/carseven/go-fullstack-template/languages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	environment := getEnvironment()
	fmt.Println("[ENVIRONMENT] " + environment)

	languages.LoadLanguagesTranslations()

	helloHandler := &handler.HelloHandler{Environment: environment}
	app.GET("/hello", helloHandler.HandleHello)
	app.GET("/hello/:name", helloHandler.HandleHello)

	mdHandle := &handler.MdHandler{Environment: environment}
	app.GET("/md", mdHandle.HandleMd)

	streamHandle := &handler.StreamHandler{Environment: environment}
	app.POST("/upload", streamHandle.HandleUpload)
	app.GET("/stream", streamHandle.HandleStream)
	app.GET("/stream/:streamId", streamHandle.HandleStream)

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
	assets := app.Group("/assets")
	assets.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// The response can be cached by browsers and intermediary caches for up to
			// 365 days (60 seconds x 60 minutes x 24 x 365 hours).
			// Immutable avoid unnecessary request to check the cache.
			// If the content is not going to change is not needed https://developer.mozilla.org/en-US/docs/Web/HTTP/Caching
			c.Response().Header().Set("Cache-Control", "max-age=31536000, immutable")
			return next(c)
		}
	})

	// TODO: Add cache-busting technic by versioning the files
	// TODO: Add generic form to access all assets (Aware to avoid leaks!!!)
	assets.File("/htmx", "assets/js/htmx@1.9.10.min.js")
	assets.File("/hls", "assets/js/hls@1.js")
	assets.File("/flowbite", "node_modules/flowbite/dist/flowbite.min.js")
	assets.File("/tailwind", "assets/css/tailwindcss/dist/style.css")
	app.Static("/video", "assets/video")

	if environment == "DEV" {
		devHandler := &handler.DevHandler{}
		app.GET("/ws", devHandler.ClientReloadWS)
		app.GET("/health", devHandler.HealthCheck)
	}

	err := app.Start(":3000")
	if err != nil {
		panic(err)
	}
}

// By default ENV will be PROD unless development specify ENV="DEV"
func getEnvironment() string {
	env := os.Getenv("ENV")

	if env != "DEV" {
		env = "PROD"
	}
	return env
}
