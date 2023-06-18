package serror

import "net/http"

// AppError represents a custom application error.
type AppError struct {
	Message    string
	StatusCode int
}

// Error returns the error message.
func (e AppError) Error() string {
	return e.Message
}

// NewAppError creates a new AppError with the specified message and status code.
func NewAppError(message string, statusCode int) AppError {
	return AppError{
		Message:    message,
		StatusCode: statusCode,
	}
}

// InternalServerError creates a new internal server error (HTTP 500) AppError.
func InternalServerError(message string) AppError {
	return NewAppError(message, http.StatusInternalServerError)
}

// BadRequestError creates a new bad request error (HTTP 400) AppError.
func BadRequestError(message string) AppError {
	return NewAppError(message, http.StatusBadRequest)
}

// NotFoundError creates a new not found error (HTTP 404) AppError.
func NotFoundError(message string) AppError {
	return NewAppError(message, http.StatusNotFound)
}

// UnauthorizedError creates a new unauthorized error (HTTP 401) AppError.
func UnauthorizedError(message string) AppError {
	return NewAppError(message, http.StatusUnauthorized)
}

// ForbiddenError creates a new forbidden error (HTTP 403) AppError.
func ForbiddenError(message string) AppError {
	return NewAppError(message, http.StatusForbidden)
}

// User alredy in Use  (HTTP 409) conflicting error AppError,
func AlreadyInUse(message string) AppError {
	return NewAppError(message, http.StatusConflict)
}
