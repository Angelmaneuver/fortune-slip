package web

import (
	"os"
	"strconv"

	"github.com/Angelmaneuver/fortune-slip/internal/lottery"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func Start(port int, lottery *lottery.Lottery) {
	var address = ":80"
	if port != 0 {
		address = ":" + strconv.Itoa(port)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(header)
	e.Use(middleware.Gzip())

	e.GET("/", func(ctx echo.Context) error {
		return ctx.File(lottery.Draw())
	})

	e.GET("/ws", func(ctx echo.Context) error {
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()

			bytes, err := os.ReadFile(lottery.Draw())
			if err != nil {
				ctx.Logger().Error(err)
			}

			err = websocket.Message.Send(ws, bytes)
			if err != nil {
				ctx.Logger().Error(err)
			}
		}).ServeHTTP(ctx.Response(), ctx.Request())

		return nil
	})

	e.Logger.Fatal(e.Start(address))
}

func header(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderCacheControl, "no-store")
		return next(ctx)
	}
}
