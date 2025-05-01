package redis

import (
	"context"
	"time"
)

type SingleClient interface {
	Set(ctx context.Context, key string, value interface{}) *StatusCmd
	SetWithExpiry(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
	MSet(ctx context.Context, keysValues ...interface{}) *StatusCmd
	Get(ctx context.Context, key string) *StringCmd
	MGet(ctx context.Context, keys ...string) *SliceCmd
	TTL(ctx context.Context, key string) *DurationCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	Exists(ctx context.Context, key string) *IntCmd
	Del(ctx context.Context, keys ...string) *IntCmd
}

func (c *client) Set(ctx context.Context, key string, value interface{}) *StatusCmd {
	return c.redisCli.Set(ctx, c.genKey(key), value, 0)
}

func (c *client) SetWithExpiry(ctx context.Context, key string, value interface{}, expiry time.Duration) *StatusCmd {
	return c.redisCli.Set(ctx, c.genKey(key), value, expiry)
}

func (c *client) MSet(ctx context.Context, keysValues ...interface{}) *StatusCmd {
	var toRedisArgs []interface{}
	for i, e := range keysValues {
		if i%2 == 0 {
			strKey, isString := e.(string)
			if !isString {
				panic("mset: key must a string")
			}
			toRedisArgs = append(toRedisArgs, c.genKey(strKey))
		} else {
			toRedisArgs = append(toRedisArgs, e)
		}
	}
	return c.redisCli.MSet(ctx, toRedisArgs...)
}

func (c *client) Get(ctx context.Context, key string) *StringCmd {
	return c.redisCli.Get(ctx, c.genKey(key))
}

func (c *client) MGet(ctx context.Context, keys ...string) *SliceCmd {
	for i, key := range keys {
		keys[i] = c.genKey(key)
	}
	return c.redisCli.MGet(ctx, keys...)
}

func (c *client) TTL(ctx context.Context, key string) *DurationCmd {
	return c.redisCli.TTL(ctx, c.genKey(key))
}

func (c *client) Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	return c.redisCli.Expire(ctx, c.genKey(key), expiration)
}

func (c *client) Exists(ctx context.Context, key string) *IntCmd {
	return c.redisCli.Exists(ctx, c.genKey(key))
}

func (c *client) Del(ctx context.Context, keys ...string) *IntCmd {
	for i, key := range keys {
		keys[i] = c.genKey(key)
	}
	return c.redisCli.Del(ctx, keys...)
}
