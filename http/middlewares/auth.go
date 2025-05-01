package middlewares

import (
	"github.com/aitu-leetcode-site/core/http/cookie"
	"github.com/aitu-leetcode-site/core/http/jwt"
	"github.com/aitu-leetcode-site/core/http/response"
	"github.com/aitu-leetcode-site/core/types"
	"github.com/valyala/fasthttp"
)

type AuthMiddleware struct{}

func (AuthMiddleware) Accept(fastHTTPH fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(requestCtx *fasthttp.RequestCtx) {
		jwtTok := cookie.ParseJWTTokenFromCookie(requestCtx)
		sessCtx, err := jwt.ParseJWT(jwtTok)
		if err != nil {
			response.SendError(requestCtx, err)
			return
		}
		requestCtx.SetUserValue(types.SessionKey, sessCtx)
		sessCtx.SetInContext(requestCtx)
		fastHTTPH(requestCtx)
	}
}
