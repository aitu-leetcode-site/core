package app

import (
	"github.com/aitu-leetcode-site/core/config"
	"github.com/aitu-leetcode-site/core/grpc"
	"github.com/aitu-leetcode-site/core/http"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

func (a *App) HttpServer() *http.Server {
	return a.httpServer
}

func (a *App) Config() *config.Config {
	return a.cfg
}

func (a *App) GrpcServer() *grpc.GRPSServer {
	return a.grpcServer
}
