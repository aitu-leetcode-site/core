package validator

import (
	"github.com/aitu-leetcode-site/core/errors"
)

type Validator interface {
	Validate() error
}

type validator struct {
	errors []error
}

func NewValidator() Validator {
	return &validator{}
}

func (v *validator) Validate() error {
	var outErr error
	for _, e := range v.errors {
		if e != nil {
			outErr = errors.Append(outErr, e)
		}
	}
	return outErr
}

func (v *validator) NotEmptyString(s string) {
	if len(s) == 0 {
		v.errors = append(v.errors, ErrValidatorEmptyString)
	}
}

func (v *validator) NotEmptyInt(i int) {
	if i == 0 {
		v.errors = append(v.errors, ErrValidatorEmptyInt)
	}
}

func (v *validator) NotEmptySlice(ss []interface{}) {
	if len(ss) == 0 {
		v.errors = append(v.errors, ErrValidatorEmptySlice)
	}
}
