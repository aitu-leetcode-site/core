package response

import (
	httpcodes "github.com/aitu-leetcode-site/core/http/codes"
	"github.com/valyala/fasthttp"
)

func SendData(requestCtx *fasthttp.RequestCtx, data interface{}) {
	baseRespData := &BaseResponse{
		StatusCode: httpcodes.StatusOK,
		Result:     data,
	}
	baseRespData.SetInCtx(requestCtx)
}
