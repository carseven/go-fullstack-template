package handler

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func GetRequestHeaderKey(c echo.Context, key string) []string {
	for headerKey, values := range c.Request().Header {
		if key == headerKey {
			fmt.Println(key)
			fmt.Println(values)
			return values
		}
	}
	return []string{}
}
