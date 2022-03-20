package rest

import (
	"fmt"
	"net/http"
)

// swagger:response errorResponse
type errorResponse struct {
	// The status code of operation
	// in: body
	Body HTTPError
}

// HTTPError is a general error returned by REST API
// swagger:model
type HTTPError struct {
	// HTTP status code
	// read only: true
	// example: 500
	Code int `json:"code"`

	// Detailed error description
	// read only: true
	// example: internal server error
	Description string `json:"description"`

	// Wrapped error
	Err error `json:"-"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Description)
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

// WithError wraps the internal error
func (e *HTTPError) WithError(err error) *HTTPError {
	e.Err = err
	return e
}

// httpError returns REST error with code and message
func httpError(code int, format string, args ...any) *HTTPError {
	return &HTTPError{
		Code:        code,
		Description: fmt.Sprintf(format, args...),
	}
}

// BadRequestErrorf returns REST error with 400 status code and message.
func BadRequestErrorf(format string, args ...any) *HTTPError {
	return httpError(http.StatusBadRequest, format, args...)
}

// InternalServerErrorf returns REST error with 500 status code and message.
func InternalServerErrorf(format string, args ...any) *HTTPError {
	return httpError(http.StatusInternalServerError, format, args...)
}
