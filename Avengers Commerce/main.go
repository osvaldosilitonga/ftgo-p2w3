package main

import (
	"net/http"
	"ngc11/config"
	"ngc11/controller"
	"ngc11/initializers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvFile()
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	db := config.InitDB()

	userController := controller.NewUserController(db)

	v1 := e.Group("/v1")
	v1.POST("/register", userController.RegisterUser)
	v1.POST("/login", userController.LoginUser)

	user := v1.Group("/user")
	user.GET("/products", func(c echo.Context) error {
		return c.String(http.StatusOK, "Products")
	})
	user.POST("/transactions", func(c echo.Context) error {
		return c.String(http.StatusOK, "Transaction")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
