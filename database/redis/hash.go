package redis

import (
	"context"
)

type HashClient interface {
	HSet(ctx context.Context, key string, field string, value interface{}) *IntCmd
	HMSet(ctx context.Context, key string, fieldsValues ...interface{}) *BoolCmd
	HGet(ctx context.Context, key string, field string) *StringCmd
	HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
	HGetAll(ctx context.Context, key string) *MapStringStringCmd
	HDel(ctx context.Context, key string, fields ...string) *IntCmd
	HExists(ctx context.Context, key, field string) *BoolCmd
}

func (c *client) HSet(ctx context.Context, key string, field string, value interface{}) *IntCmd {
	return c.redisCli.HSet(ctx, c.genKey(key), field, value)
}

func (c *client) HMSet(ctx context.Context, key string, fieldsValues ...interface{}) *BoolCmd {
	return c.redisCli.HMSet(ctx, c.genKey(key), fieldsValues...)
}

func (c *client) HGet(ctx context.Context, key string, field string) *StringCmd {
	return c.redisCli.HGet(ctx, c.genKey(key), field)
}

func (c *client) HMGet(ctx context.Context, key string, fields ...string) *SliceCmd {
	return c.redisCli.HMGet(ctx, c.genKey(key), fields...)
}

func (c *client) HGetAll(ctx context.Context, key string) *MapStringStringCmd {
	return c.redisCli.HGetAll(ctx, c.genKey(key))
}

func (c *client) HDel(ctx context.Context, key string, fields ...string) *IntCmd {
	return c.redisCli.HDel(ctx, c.genKey(key), fields...)
}

func (c *client) HExists(ctx context.Context, key, field string) *BoolCmd {
	return c.redisCli.HExists(ctx, c.genKey(key), field)
}
