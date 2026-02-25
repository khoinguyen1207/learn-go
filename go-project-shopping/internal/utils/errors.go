package utils

type AppError struct {
	Code    ErrorResponseCode
	Message string
	Err     error
}

type ErrorResponseCode string

const (
	CodeBadRequest          ErrorResponseCode = "BAD_REQUEST"
	CodeUnauthorized        ErrorResponseCode = "UNAUTHORIZED"
	CodeForbidden           ErrorResponseCode = "FORBIDDEN"
	CodeNotFound            ErrorResponseCode = "NOT_FOUND"
	CodeConflict            ErrorResponseCode = "CONFLICT"
	CodeInternalServerError ErrorResponseCode = "INTERNAL_SERVER_ERROR"
	CodeValidationError     ErrorResponseCode = "VALIDATION_ERROR"
	CodeTooManyRequests     ErrorResponseCode = "TOO_MANY_REQUESTS"
)

func (e *AppError) Error() string {
	return ""
}

func NewError(message string, code ErrorResponseCode) error {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func WrapError(err error, message string, code ErrorResponseCode) error {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
