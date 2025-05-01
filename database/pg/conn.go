package pg

import (
	"context"
	"github.com/aitu-leetcode-site/core/database/pg/config"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type conn struct {
	noCopy noCopy
	pgConn *pgx.Conn
}

func NewConn(pgConfig *pgconfig.Config) (Conn, error) {
	return &conn{}, nil
}

func (c *conn) Exec(ctx context.Context, query string, args ...interface{}) (int64, error) {
	cmdTag, err := c.pgConn.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return cmdTag.RowsAffected(), nil
}

func (c *conn) ScanRow(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, c.pgConn, dst, query, args...)
}

func (c *conn) Scan(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, c.pgConn, dst, query, args...)
}

func (c *conn) Close(ctx context.Context) error {
	return c.pgConn.Close(ctx)
}
