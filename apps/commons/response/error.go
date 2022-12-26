package response

import "net/http"

type CustomError struct {
	StatusCode     int         `json:"status_code"`
	Message        string      `json:"message"`
	AdditionalInfo interface{} `json:"additional_info,omitempty"`
}

var (
	generalError = CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
	}
	repositoryError = CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    "repository error",
	}
	notFoundError = CustomError{
		StatusCode: http.StatusNotFound,
		Message:    "not found error",
	}
	unauthorizedError = CustomError{
		StatusCode: http.StatusUnauthorized,
		Message:    "unauthorized",
	}
)

func GeneralError() *CustomError {
	err := generalError
	return &err
}
func GeneralErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := generalError
	err.AdditionalInfo = info
	return &err
}

func RepositoryError() *CustomError {
	err := repositoryError
	return &err
}

func RepositoryErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := repositoryError
	err.AdditionalInfo = info
	return &err
}
func NotFoundError() *CustomError {
	err := notFoundError
	return &err
}

func NotFoundErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := repositoryError
	err.AdditionalInfo = info
	return &err
}
func UnauthorizedError() *CustomError {
	err := unauthorizedError
	return &err
}

func UnauthorizedErrorWithAdditionalInfo(info interface{}) *CustomError {
	err := unauthorizedError
	err.AdditionalInfo = info
	return &err
}
