package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! from ping.")
	})
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! from hello.")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
