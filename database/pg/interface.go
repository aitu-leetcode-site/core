package pg

import (
	"context"
	"github.com/aitu-leetcode-site/core/app/component"
)

type Conn interface {
	component.Closer
	Exec(ctx context.Context, query string, args ...interface{}) (int64, error)
	ScanRow(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	Scan(ctx context.Context, dst interface{}, query string, args ...interface{}) error
}
