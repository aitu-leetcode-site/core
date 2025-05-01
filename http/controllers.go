package http

import fasthttprouter "github.com/fasthttp/router"

type Controller interface {
	Init(fasthttpR *fasthttprouter.Router)
}

func (s *Server) WithControllers(cntrls ...Controller) {
	for _, ctrl := range cntrls {
		ctrl.Init(s.router)
	}
}
