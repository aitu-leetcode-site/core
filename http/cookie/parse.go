package cookie

import (
	"github.com/valyala/fasthttp"
)

const (
	cookieName = "backend_auth"
)

func ParseJWTTokenFromCookie(requestCtx *fasthttp.RequestCtx) string {
	cookieBytes := requestCtx.Request.Header.Cookie(cookieName)
	return string(cookieBytes)
}
