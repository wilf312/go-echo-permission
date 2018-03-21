package main

import (
	"net/http"
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "./handler"
)

func main() {
	const port = ":1323"

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", route.MainPage())
  // e.GET("/path/:pathParameter", handler.PathParameter()
)

	// Start server
	e.Logger.Fatal(e.Start(port))
}
