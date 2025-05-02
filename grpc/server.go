package grpc

import (
	"context"
	"github.com/aitu-leetcode-site/core/log"
	"github.com/aitu-leetcode-site/core/utils/port"
	"google.golang.org/grpc"
	"net"
	"time"
)

type GRPSServer struct {
	*grpc.Server
	startErr error
	port     port.Port
}

func NewGRPCServer() *GRPSServer {
	grpsServer := &GRPSServer{
		port:   gRPCPort,
		Server: grpc.NewServer(),
	}
	return grpsServer
}

func (g *GRPSServer) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", newGRPCAddress(port.AnyAddress))
	if err != nil {
		return err
	}
	log.Infof(ctx, "GRPC Server start listen at port %s", g.port.String())
	go g.serve(listener)
	time.Sleep(time.Millisecond * 100)
	if g.startErr != nil {
		return err
	}
	return nil
}

func (g *GRPSServer) serve(listener net.Listener) {
	err := g.Server.Serve(listener)
	if err != nil {
		g.startErr = err
	}
	return
}

func (g *GRPSServer) Close(_ context.Context) error {
	g.Server.Stop()
	return nil
}
