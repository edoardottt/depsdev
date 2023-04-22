package client

import (
	"errors"
	"fmt"
)

var (
	ErrStatus = errors.New("server replied with status: ")
)

type ErrorResponse struct {
	Message string `json:"error,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}
