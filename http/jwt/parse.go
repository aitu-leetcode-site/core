package jwt

import (
	"github.com/aitu-leetcode-site/core/errors"
	"github.com/aitu-leetcode-site/core/types"
	"github.com/golang-jwt/jwt"
)

func ParseJWT(jwtS string) (*types.Session, error) {
	jwtToken, err := jwt.Parse(jwtS, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrJWTOldVersion
		}
		return []byte(signature), nil
	})
	if jwtToken == nil {
		jwtToken = &jwt.Token{}
	}
	if !jwtToken.Valid {
		switch {
		case errors.Is(err, jwt.ErrSignatureInvalid):
			return nil, errors.ErrJWTSignatureInvalid
		}
		return nil, errors.ErrJWTOldVersion
	}
	if _, ok := jwtToken.Claims.(jwt.MapClaims); !ok {
		return nil, errors.ErrJWTOldVersion
	}
	jwtClaims := parseFromJWTClaims(jwtToken.Claims.(jwt.MapClaims))
	err = jwtClaims.Validate()
	if err != nil {
		return nil, err
	}
	sess := types.NewSession(
		jwtClaims.findByName(userIDClaim).value.(types.UserID),
	)

	return sess, nil
}
