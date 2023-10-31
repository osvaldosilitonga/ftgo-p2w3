package main

import (
	"net/http"
	"ngc11/config"
	"ngc11/initializers"

	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvFile()

}

func main() {
	app := echo.New()
	db := config.InitDB()

	v1 := app.Group("/v1")
	v1.POST("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "Login")
	})
	v1.POST("/register", func(c echo.Context) error {
		return c.String(http.StatusOK, "Register")
	})

	user := v1.Group("/user")
	user.GET("/products", func(c echo.Context) error {
		return c.String(http.StatusOK, "Products")
	})
	user.POST("/transactions", func(c echo.Context) error {
		return c.String(http.StatusOK, "Transaction")
	})

	app.Logger.Fatal(app.Start(":8080"))
}
