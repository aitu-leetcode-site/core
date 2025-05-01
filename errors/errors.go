package errors

import (
	"errors"
	"github.com/aitu-leetcode-site/core/http/codes"
	"github.com/hashicorp/go-multierror"
)

var (
	As     = errors.As
	Is     = errors.Is
	New    = errors.New
	Join   = errors.Join
	Unwrap = errors.Unwrap
)

var (
	Append         = multierror.Append
	Prefix         = multierror.Prefix
	ListFormatFunc = multierror.ListFormatFunc
)

var (
	ErrAuthRequired = &ApiError{HTTPCode: httpcodes.StatusUnauthorized, Message: "Authorization is required"}
)

var (
	ErrNotFound = &ApiError{HTTPCode: httpcodes.StatusNotFound, Message: "Not found"}
)

var (
	ErrJWTSignatureInvalid = &ApiError{HTTPCode: httpcodes.StatusUnauthorized, Message: "JWT signature is invalid"}
	ErrJWTInvalidClaims    = &ApiError{HTTPCode: httpcodes.StatusBadRequest, Message: "JWT claims is invalid"}
	ErrJWTExpired          = &ApiError{HTTPCode: httpcodes.StatusUnauthorized, Message: "JWT expired"}
	ErrJWTOldVersion       = &ApiError{HTTPCode: httpcodes.StatusUnauthorized, Message: "JWT version is old"}
)
