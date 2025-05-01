package redis

import "context"

type SetClient interface {
	SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd
	SRem(ctx context.Context, key string, members ...interface{}) *IntCmd
	SCard(ctx context.Context, key string) *IntCmd
	SRandMember(ctx context.Context, key string) *StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd
	SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd
}

func (c *client) SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd {
	return c.redisCli.SAdd(ctx, c.genKey(key), members...)
}

func (c *client) SRem(ctx context.Context, key string, members ...interface{}) *IntCmd {
	return c.redisCli.SRem(ctx, c.genKey(key), members...)
}

func (c *client) SCard(ctx context.Context, key string) *IntCmd {
	return c.redisCli.SCard(ctx, c.genKey(key))
}

func (c *client) SRandMember(ctx context.Context, key string) *StringCmd {
	return c.redisCli.SRandMember(ctx, c.genKey(key))
}

func (c *client) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd {
	return c.redisCli.SRandMemberN(ctx, c.genKey(key), count)
}

func (c *client) SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd {
	return c.redisCli.SIsMember(ctx, c.genKey(key), member)
}
