package http

import "github.com/aitu-leetcode-site/core/http/middlewares"

func (s *Server) WithMiddlewares(ms ...middlewares.Middleware) {
	s.middlewares = append(s.middlewares, ms...)
}

func (s *Server) acceptMiddlewares() {
	handler := s.router.Handler
	for _, middleware := range s.middlewares {
		handler = middleware.Accept(handler)
	}
	s.fasthttpSrv.Handler = handler
}
