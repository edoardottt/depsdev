package depsdev

import (
	"fmt"
)

type ErrorResponse struct {
	Message string `json:"error,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
