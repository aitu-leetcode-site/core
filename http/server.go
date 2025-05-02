package http

import (
	"context"
	httpconfig "github.com/aitu-leetcode-site/core/http/config"
	"github.com/aitu-leetcode-site/core/http/middlewares"
	"github.com/aitu-leetcode-site/core/log"
	"github.com/aitu-leetcode-site/core/utils/port"
	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"time"
)

type Server struct {
	fasthttpSrv *fasthttp.Server
	chans       struct {
		isStartedChan chan struct{}
	}
	startErr    error
	port        port.Port
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
		chans: struct {
			isStartedChan chan struct{}
		}{isStartedChan: make(chan struct{})},
	}
	corsM := middlewares.NewCORSMiddleware()
	loggingMiddleware := middlewares.NewLoggingMiddleware()
	outServer.WithMiddlewares(corsM, loggingMiddleware)
	return outServer
}

func (s *Server) Start(ctx context.Context) error {
	s.acceptMiddlewares()
	go s.serve()
	log.Infof(ctx, "HTTP Server start listen at port %v", s.port.String())
	time.Sleep(time.Millisecond * 100)
	if s.startErr != nil {
		return s.startErr
	}
	return nil
}

func (s *Server) serve() {
	addr := s.port.Addr(port.AnyAddress)
	switch err := s.fasthttpSrv.ListenAndServe(addr); {
	case err == nil:
		return
	default:
		s.startErr = err
	}
	return
}

func (s *Server) Close(_ context.Context) error {
	return s.fasthttpSrv.Shutdown()
}
