package utils

import "fmt"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *AppError) GetMessage() string {
	return e.Message
}

func (e *AppError) GetCode() int {
	return e.Code
}

func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}
