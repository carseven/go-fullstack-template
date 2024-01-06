package handler

import (
	"github.com/carseven/go-fullstack-template/view/cases/hello"
	"github.com/carseven/go-fullstack-template/view/components/footer"
	"github.com/carseven/go-fullstack-template/view/components/navbar"
	layout "github.com/carseven/go-fullstack-template/view/layouts"
	"github.com/labstack/echo/v4"
)

type HelloHandler struct {
	Environment string
}

func (h HelloHandler) HandleHello(c echo.Context) error {
	name := c.Param("name")

	userLanguage := GetRequestHeaderKey(c, "Accept-Language")
	return render(c, layout.Base(hello.View(name, userLanguage), navbar.NavbarComponent(), footer.FooterComponent(), h.Environment))
}
