package middlewares

import (
	"github.com/valyala/fasthttp"
)

type CORSMiddleware struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

func NewCORSMiddleware() Middleware {
	return &CORSMiddleware{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}
}

func (c *CORSMiddleware) Accept(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		origin := string(ctx.Request.Header.Peek("Origin"))
		if origin == "" {
			next(ctx)
			return
		}

		ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
		ctx.Response.Header.Set("Vary", "Origin")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", join(c.AllowedMethods))
		ctx.Response.Header.Set("Access-Control-Allow-Headers", join(c.AllowedHeaders))
		if c.AllowCredentials {
			ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		}
		if string(ctx.Method()) == fasthttp.MethodOptions {
			ctx.SetStatusCode(fasthttp.StatusNoContent)
			return
		}
		next(ctx)
	}
}

func join(items []string) string {
	var out []byte
	for i, item := range items {
		if i > 0 {
			out = append(out, ',')
			out = append(out, ' ')
		}
		out = append(out, item...)
	}
	return string(out)
}
