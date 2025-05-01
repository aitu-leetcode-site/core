package validator

import (
	"github.com/aitu-leetcode-site/core/errors"
	"github.com/aitu-leetcode-site/core/http/codes"
)

var (
	ErrValidatorEmptyString = errors.NewApiError(
		httpcodes.StatusBadRequest,
		"Empty string",
	)
	ErrValidatorEmptyInt = errors.NewApiError(
		httpcodes.StatusBadRequest,
		"Empty integer",
	)
	ErrValidatorEmptySlice = errors.NewApiError(
		httpcodes.StatusBadRequest,
		"Empty slice",
	)
)
