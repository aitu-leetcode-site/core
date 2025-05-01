package middlewares

import (
	"fmt"
	"github.com/aitu-leetcode-site/core/errors"
	codes "github.com/aitu-leetcode-site/core/http/codes"
	"github.com/aitu-leetcode-site/core/http/response"
	"github.com/valyala/fasthttp"
)

type RecoverMiddleware struct{}

func (ms *RecoverMiddleware) Accept(fastHTTPH fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if r := recover(); r != nil {
				response.SendError(ctx, errors.NewApiError(
					codes.StatusServiceUnavailable, fmt.Sprint(r)),
				)
			}
		}()
		fastHTTPH(ctx)
	}
}
