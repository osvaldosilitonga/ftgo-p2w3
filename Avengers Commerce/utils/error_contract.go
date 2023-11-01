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
	ErrInvalidParamID = APIError{
		Code:    http.StatusBadRequest,
		Message: "Invalid Param ID",
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
	ErrUnauthorized = APIError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
	ErrTokenAlg = APIError{
		Code:    http.StatusUnauthorized,
		Message: "Unexpected signing method",
	}
	ErrInvalidToken = APIError{
		Code:    http.StatusUnauthorized,
		Message: "Invalid Access Token",
	}
	ErrTokenExpired = APIError{
		Code:    http.StatusUnauthorized,
		Message: "Token expired",
	}
)
