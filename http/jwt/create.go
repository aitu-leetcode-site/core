package jwt

import (
	"github.com/aitu-leetcode-site/core/types"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	jwtExpired     = time.Hour * 24
	refreshExpired = time.Hour * 24 * 30
)

func GenerateJWT(sess *types.Session, isRefresh bool) (string, error) {
	expiration := time.Now().Add(jwtExpired)
	if isRefresh {
		expiration = time.Now().Add(refreshExpired)
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDClaim:  sess.UserID(),
		refreshClaim: isRefresh,
		expClaim:     expiration,
	})
	signed, err := jwtToken.SignedString([]byte(signature))
	if err != nil {
		return "", err
	}
	return signed, nil
}
