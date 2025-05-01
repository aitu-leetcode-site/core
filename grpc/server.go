package grpc

import (
	"context"
	"github.com/aitu-leetcode-site/core/utils"
	"google.golang.org/grpc"
	"net"
)

type GRPSServer struct {
	*grpc.Server
	port utils.Port
}

func NewGRPCServer() *GRPSServer {
	grpsServer := &GRPSServer{
		port:   utils.Port(8080),
		Server: grpc.NewServer(),
	}
	return grpsServer
}

func (g *GRPSServer) Start(context.Context) error {
	listener, err := net.Listen("tcp", g.port.Addr("0.0.0.0"))
	if err != nil {
		return err
	}
	return g.Server.Serve(listener)
}

func (g *GRPSServer) Close(_ context.Context) error {
	g.Server.Stop()
	return nil
}
