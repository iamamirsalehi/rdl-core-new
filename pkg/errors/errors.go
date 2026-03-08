package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func Wrap(code int, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}

var (
	ErrNotFound     = New(http.StatusNotFound, "resource not found")
	ErrUnauthorized = New(http.StatusUnauthorized, "unauthorized")
	ErrForbidden    = New(http.StatusForbidden, "forbidden")
	ErrBadRequest   = New(http.StatusBadRequest, "bad request")
	ErrInternal     = New(http.StatusInternalServerError, "internal server error")
	ErrConflict     = New(http.StatusConflict, "resource already exists")
)

func IsNotFound(err error) bool {
	var e *AppError
	return errors.As(err, &e) && e.Code == http.StatusNotFound
}
