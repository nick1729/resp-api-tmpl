package errors

import (
	"fmt"
)

type ExternalError struct {
	Code    int    `example:"1"`
	Message string `example:"Unable to parse request"`
	Err     string `example:"Extended description of error" json:"details,omitempty"`
}

func (e *ExternalError) Error() string {
	return fmt.Sprintf("error code: %d; error message: %s; details: %s", e.Code, e.Message, e.Err)
}

func NewFromBase(base ExternalError, details error) *ExternalError {
	return &ExternalError{
		Code:    base.Code,
		Message: base.Message,
		Err:     details.Error(),
	}
}

func New(code int, message string) *ExternalError {
	return &ExternalError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrParseRequest = ExternalError{
		Code:    1,
		Message: "unable to parse request",
	}

	ErrValidateRequest = ExternalError{
		Code:    2,
		Message: "unable to validate request",
	}
)
