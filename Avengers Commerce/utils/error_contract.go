package utils

import "net/http"

type APIError struct {
	Code    int
	Message string
}

var (
	// 200
	ErrDataNotFound = APIError{
		Code:    http.StatusOK,
		Message: "Empty Data",
	}

	// 400
	ErrBadRequest = APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
	}
	ErrDuplicateKey = APIError{
		Code:    http.StatusBadRequest,
		Message: "Duplicate key, unique constraint",
	}
	ErrUsernameNotFound = APIError{
		Code:    http.StatusBadRequest,
		Message: "Username not exist",
	}
	ErrInvalidPassword = APIError{
		Code:    http.StatusBadRequest,
		Message: "Invalid Password",
	}
	ErrNotFound = APIError{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}

	// 500
	ErrInternalServer = APIError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	ErrGenerateToken = APIError{
		Code:    http.StatusInternalServerError,
		Message: "Failed when generate token",
	}
)
