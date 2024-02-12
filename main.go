package main

import (
	"hello-echo/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlers.Hello)
	e.GET("/hello/:name", handlers.HelloName)
	e.GET("/ping", handlers.Pong)

	// Start server
	e.Logger.Fatal(e.Start(":8008"))
}
