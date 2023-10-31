package controller

import (
	"errors"
	"net/http"
	"ngc11/dto"
	"ngc11/entity"
	"ngc11/repository"
	"ngc11/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) User {
	return User{
		DB: db,
	}
}

func (controller User) RegisterUser(c echo.Context) error {
	body := dto.RegisterUserBody{}

	c.Bind(&body)
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ErrBadRequest)
	}

	// Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	user := entity.Users{
		Username:      body.Username,
		Password:      string(hash),
		DepositAmount: body.DepositAmount,
	}

	// Create to DB
	if err = repository.CreateUser(&user, controller.DB); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return utils.ErrorMessage(c, &utils.ErrDuplicateKey)
		}
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	return c.JSON(http.StatusCreated, user)
}

func (controller User) LoginUser(c echo.Context) error {
	body := dto.LoginUserBody{}

	c.Bind(&body)
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ErrBadRequest)
	}

	// get user
	user, err := repository.GetUserByUsername(body.Username, controller.DB)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrorMessage(c, &utils.ErrUsernameNotFound)
		}
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return utils.ErrorMessage(c, &utils.ErrInvalidPassword)
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 5).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
	if err != nil {
		return utils.ErrorMessage(c, &utils.ErrGenerateToken)
	}

	response := dto.LoginResponse{
		Code:        http.StatusOK,
		Message:     "Login Success",
		AccessToken: tokenString,
	}

	return c.JSON(http.StatusOK, response)
}
