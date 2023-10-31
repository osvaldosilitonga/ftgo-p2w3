package middleware

import (
	"errors"
	"ngc11/repository"
	"ngc11/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

func NewAuthHandler(db *gorm.DB) Auth {
	return Auth{
		DB: db,
	}
}

func (handler Auth) RequiredAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get token from header
		authHeader := c.Request().Header.Get("authorization")

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, utils.ErrorMessage(c, &utils.ErrTokenAlg)
			}

			return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
		})
		if err != nil {
			return utils.ErrorMessage(c, &utils.ErrInvalidToken)
		}

		// Token Claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the exp date
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				return utils.ErrorMessage(c, &utils.ErrTokenExpired)
			}

			// Find the user with token
			user, err := repository.GetUserByUsername(claims["username"].(string), handler.DB)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.ErrorMessage(c, &utils.ErrUnauthorized)
			}
			if err != nil {
				return utils.ErrorMessage(c, &utils.ErrInternalServer)
			}

			// Attach to req
			c.Set("user", user.Username)

			// Continue
			return next(c)

		} else {
			return utils.ErrorMessage(c, &utils.ErrInvalidToken)
		}

	}
}
