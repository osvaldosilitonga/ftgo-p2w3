package main

import (
	"net/http"
	"ngc11/config"
	"ngc11/controller"
	"ngc11/initializers"
	"ngc11/middleware"

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

	authMiddleware := middleware.NewAuthHandler(db)

	userController := controller.NewUserController(db)
	storeController := controller.NewStoreController(db)

	v1 := e.Group("/v1")
	v1.POST("/register", userController.RegisterUser)
	v1.POST("/login", userController.LoginUser)

	user := v1.Group("/user")
	user.Use(authMiddleware.RequiredAuth)
	user.GET("/products", userController.GetProducts)
	user.POST("/transactions", userController.CreateTransaction)
	user.GET("/stores", storeController.GetAllStore)
	user.GET("/stores/:id", storeController.GetStoreDetailByID)

	e.Logger.Fatal(e.Start(":8080"))
}
