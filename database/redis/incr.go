package redis

import "context"

type IncrClient interface {
	Incr(ctx context.Context, key string) *IntCmd
	IncrBy(ctx context.Context, key string, incrBy int64) *IntCmd
	Decr(ctx context.Context, key string) *IntCmd
	DecrBy(ctx context.Context, key string, incrBy int64) *IntCmd
	DecrByFloat(ctx context.Context, key string, incrBy float64) *FloatCmd
	IncrByFloat(ctx context.Context, key string, incrBy float64) *FloatCmd
	HIncrBy(ctx context.Context, key string, hIncrBy int64) *IntCmd
	HIncrByFloat(ctx context.Context, key string, field string, hIncrBy float64) *FloatCmd
}

func (c *client) Incr(ctx context.Context, key string) *IntCmd {
	return c.redisCli.Incr(ctx, c.genKey(key))
}

func (c *client) IncrBy(ctx context.Context, key string, incrBy int64) *IntCmd {
	return c.redisCli.IncrBy(ctx, c.genKey(key), incrBy)
}

func (c *client) Decr(ctx context.Context, key string) *IntCmd {
	return c.redisCli.Decr(ctx, c.genKey(key))
}

func (c *client) DecrBy(ctx context.Context, key string, incrBy int64) *IntCmd {
	return c.redisCli.DecrBy(ctx, c.genKey(key), incrBy)
}

func (c *client) DecrByFloat(ctx context.Context, key string, incrBy float64) *FloatCmd {
	return c.redisCli.IncrByFloat(ctx, c.genKey(key), -1*incrBy)
}

func (c *client) IncrByFloat(ctx context.Context, key string, incrBy float64) *FloatCmd {
	return c.redisCli.IncrByFloat(ctx, c.genKey(key), incrBy)
}

func (c *client) HIncrBy(ctx context.Context, key string, hIncrBy int64) *IntCmd {
	return c.redisCli.IncrBy(ctx, c.genKey(key), hIncrBy)
}

func (c *client) HIncrByFloat(ctx context.Context, key string, field string, hIncrBy float64) *FloatCmd {
	return c.redisCli.HIncrByFloat(ctx, c.genKey(key), field, hIncrBy)
}
