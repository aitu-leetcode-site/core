package grpc

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

type GRPCClient struct {
	conn *grpc.ClientConn
}

func NewClient(ctx context.Context, addr string) (*GRPCClient, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	conn, err := grpc.DialContext(ctx, addr)
	if err != nil {
		return nil, err
	}
	return &GRPCClient{
		conn: conn,
	}, nil
}

func (c *GRPCClient) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	return c.conn.Invoke(ctx, method, args, reply, opts...)
}

func (c *GRPCClient) NewStream(
	ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	return c.conn.NewStream(ctx, desc, method, opts...)
}

func (c *GRPCClient) Close(_ context.Context) error {
	return c.conn.Close()
}
