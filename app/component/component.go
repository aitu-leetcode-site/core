package component

import "context"

type Starter interface {
	Start(ctx context.Context) error
}

type Closer interface {
	Close(ctx context.Context) error
}

type StarterCloser interface {
	Starter
	Closer
}
