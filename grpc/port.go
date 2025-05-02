package grpc

import "github.com/aitu-leetcode-site/core/utils/port"

const gRPCPort = 3009

func newGRPCPort() *port.Port {
	var outPort = port.Port(gRPCPort)
	return &outPort
}

func newGRPCAddress(addr string) string {
	return newGRPCPort().Addr(addr)
}
