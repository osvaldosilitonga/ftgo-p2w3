package main

import (
	"net/http"
	"ngc11/initializers"

	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvFile()
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
