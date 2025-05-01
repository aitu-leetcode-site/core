package http

import (
	"context"
	httpconfig "github.com/aitu-leetcode-site/core/http/config"
	"github.com/aitu-leetcode-site/core/http/middlewares"
	"github.com/aitu-leetcode-site/core/utils"
	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Server struct {
	fasthttpSrv *fasthttp.Server
	port        utils.Port
	middlewares []middlewares.Middleware
	router      *fasthttprouter.Router
}

func NewServer(appName string, config *httpconfig.Config) *Server {
	r := fasthttprouter.New()
	fasthttpS := &fasthttp.Server{
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		Handler:      r.Handler,
		Name:         appName,
	}
	outServer := &Server{
		fasthttpSrv: fasthttpS,
		port:        config.Port,
		router:      r,
	}
	corsM := middlewares.NewCORSMiddleware()
	loggingMiddleware := middlewares.NewLoggingMiddleware()
	outServer.WithMiddlewares(corsM, loggingMiddleware)
	return outServer
}

func (s *Server) Start(_ context.Context) error {
	return s.fasthttpSrv.ListenAndServe("0.0.0.0:" + s.port.String())
}

func (s *Server) Close(ctx context.Context) error {
	return s.fasthttpSrv.Shutdown()
}
