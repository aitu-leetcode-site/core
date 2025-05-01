package middlewares

import "github.com/valyala/fasthttp"

type Middleware interface {
	Accept(fastHTTPH fasthttp.RequestHandler) fasthttp.RequestHandler
}
