package middlewares

import (
	"github.com/aitu-leetcode-site/core/log"
	realip "github.com/ferluci/fast-realip"

	"github.com/valyala/fasthttp"
)

type LoggingMiddleware struct{}

func (l LoggingMiddleware) Accept(fastHTTPH fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(requestCtx *fasthttp.RequestCtx) {
		logger := log.With(
			log.NewField("method", string(requestCtx.Method())),
			log.NewField("path", string(requestCtx.Path())),
			log.NewField("ip", realip.FromRequest(requestCtx)),
		)
		logger.Infof(requestCtx, "Got HTTP request")
		fastHTTPH(requestCtx)
	}
}

func NewLoggingMiddleware() Middleware {
	return &LoggingMiddleware{}
}
