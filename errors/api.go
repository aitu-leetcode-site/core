package errors

import (
	"fmt"
	"github.com/aitu-leetcode-site/core/http/codes"
)

type ApiError struct {
	HTTPCode httpcodes.HTTPCode
	Message  string
}

func NewApiError(HTTPCode httpcodes.HTTPCode, message string) *ApiError {
	return &ApiError{HTTPCode: HTTPCode, Message: message}
}

func (e ApiError) StatusCode() int {
	return e.HTTPCode.StatusCode()
}

func NewDefaultUserApiError(message string) *ApiError {
	return &ApiError{
		HTTPCode: httpcodes.StatusBadRequest,
		Message:  message,
	}
}

func NewBadServerApiError(message string) *ApiError {
	return &ApiError{
		HTTPCode: httpcodes.StatusInternalServerError,
		Message:  message,
	}
}

func (e ApiError) Error() string {
	if len(e.Message) == 0 {
		return ""
	}
	if e.HTTPCode == 0 {
		return e.Message
	}
	return fmt.Sprintf("%d : %s", e.HTTPCode, e.Message)
}

func (e ApiError) String() string {
	return e.Message
}
