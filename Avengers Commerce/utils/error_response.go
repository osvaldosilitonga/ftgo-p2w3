package utils

import "github.com/labstack/echo/v4"

func ErrorMessage(c echo.Context, apiError *APIError) error {
	return c.JSON(
		apiError.Code,
		apiError,
	)
}
