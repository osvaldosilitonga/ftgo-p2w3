package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Register struct {
	DB *gorm.DB
}

func NewRegisterHandler(db *gorm.DB) Register {
	return Register{
		DB: db,
	}
}

func (controller Register) RegisterUser(c echo.Context) error {
	return c.String(http.StatusOK, "Register User")
}
