package handlers

import (
	"hello-echo/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler
func HelloName(c echo.Context) error {
	title := "Hello:" + c.Param("name")
	name := c.Param("name")
	//return templates.Index(title, name).Render(context.Background(), c.Response().Writer)
	return render(c, view.Index(title, name))
}

func Pong(c echo.Context) error {
	//return templates.Index("Pong", "FuncPong").Render(context.Background(), c.Response().Writer)
	return render(c, view.Index("Pong", "FuncPong"))
}
