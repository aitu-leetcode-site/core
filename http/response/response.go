package response

import (
	httpcodes "github.com/aitu-leetcode-site/core/http/codes"
	"github.com/valyala/fasthttp"
)

//go:generate easyjson -all response.go

type BaseResponse struct {
	StatusCode    httpcodes.HTTPCode `json:"status_code"`
	ErrorsMessage *string            `json:"errors,omitempty"`
	Result        interface{}        `json:"result,omitempty"`
}

func NewBaseResponse(statusCode httpcodes.HTTPCode) *BaseResponse {
	return &BaseResponse{
		StatusCode: statusCode,
	}
}

func (r *BaseResponse) SetInCtx(requestCtx *fasthttp.RequestCtx) {
	bytes, err := r.MarshalJSON()
	if err != nil {
		requestCtx.Response.SetStatusCode(httpcodes.StatusNotExtended.StatusCode())
		return
	}
	requestCtx.Response.SetBody(bytes)
}
