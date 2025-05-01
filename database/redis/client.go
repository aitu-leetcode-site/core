package redis

import (
	"fmt"
	"github.com/aitu-leetcode-site/core/database/redis/config"
	"github.com/redis/go-redis/v9"
)

type Client interface {
	SingleClient
	HashClient
	JsonClient
	SetClient
	IncrClient
}

type client struct {
	appPrefix string
	redisCli  redis.UniversalClient
}

func NewClient(config *redisconfig.Config) Client {
	return &client{}
}

const clientKeyFormat = "%s:%s"

func (c *client) genKey(key string) string {
	return fmt.Sprintf(clientKeyFormat, c.appPrefix, key)
}
