package jwt

import (
	"github.com/aitu-leetcode-site/core/errors"
	"github.com/aitu-leetcode-site/core/types"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type claim struct {
	field string
	value any
}

type claims map[string]claim

func parseFromJWTClaims(mapClaims jwt.MapClaims) claims {
	var out = make(claims)
	for key, value := range mapClaims {
		out[key] = claim{field: key, value: value}
	}
	return out
}

var (
	expClaim     = "exp"
	userIDClaim  = "userID"
	refreshClaim = "refresh"
)

func (cl claims) Validate() error {
	var outE error
	for _, c := range cl {
		switch c.field {
		case expClaim:
			if c.value == nil {
				outE = errors.Append(outE, errors.ErrJWTInvalidClaims)
				break
			}
			expTime, expTimeOk := c.value.(time.Time)
			if !expTimeOk {
				outE = errors.Append(outE, errors.ErrJWTInvalidClaims)
				break
			}
			if expTime.Before(time.Now()) {
				outE = errors.Append(outE, errors.ErrJWTExpired)
				break
			}
		case userIDClaim:
			if c.value == nil {
				outE = errors.Append(outE, errors.ErrJWTInvalidClaims)
				break
			}
			userID, userIDOk := c.value.(types.UserID)
			if !userIDOk {
				outE = errors.Append(outE, errors.ErrJWTInvalidClaims)
				break
			}
			err := userID.Validate()
			if err != nil {
				outE = errors.Append(outE, errors.ErrJWTInvalidClaims)
				break
			}
		case refreshClaim:
			if c.value == nil {
				outE = errors.Append(outE, errors.ErrJWTInvalidClaims)

			}
			isRefresh, isRefreshOk := c.value.(bool)
			if !isRefreshOk {
				outE = errors.Append(outE, errors.ErrJWTInvalidClaims)
				break
			}
			if isRefresh {
				outE = errors.Append(outE, errors.ErrJWTInvalidClaims)
				break
			}
		}
	}
	return outE
}

func (cl claims) findByName(name string) claim {
	return cl[strings.ToLower(name)]
}
