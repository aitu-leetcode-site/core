package response

import (
	"github.com/aitu-leetcode-site/core/errors"
	"github.com/valyala/fasthttp"
)

func SendError(requestCtx *fasthttp.RequestCtx, err error) {
	var apiError errors.ApiError
	if errors.As(err, &apiError) {
		statusCode := apiError.HTTPCode
		requestCtx.SetStatusCode(apiError.StatusCode())
		outResp := NewBaseResponse(statusCode)
		outResp.ErrorsMessage = &apiError.Message
		outResp.SetInCtx(requestCtx)
		return
	}
	outResp := NewBaseResponse(fasthttp.StatusBadRequest)
	errMessage := err.Error()
	outResp.ErrorsMessage = &errMessage
	outResp.SetInCtx(requestCtx)
}
