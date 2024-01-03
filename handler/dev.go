package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type DevHandler struct{}

// Websocket to indicate the client when the air server reloads.
// When the server WS connection close because of a rebuild, the client will poll the health check when the server build
// has finished and reload the client
func (h DevHandler) ClientReloadWS(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		err := websocket.Message.Send(ws, "CONNECTED")
		if err != nil {
			c.Logger().Error(err)
		}
		for {
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				break
			}

			if msg == "PING" {
				c.Logger().Info(msg)
				err := websocket.Message.Send(ws, "PONG")
				if err != nil {
					c.Logger().Error(err)
					break
				}
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func (h DevHandler) HealthCheck(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.NoContent(http.StatusOK)
}
