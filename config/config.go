package config

import (
	pgconfig "github.com/aitu-leetcode-site/core/database/pg/config"
	refisconfig "github.com/aitu-leetcode-site/core/database/redis/config"
)

type Config struct {
	noCopy      noCopy
	AppName     string
	PgConfig    *pgconfig.Config
	RedisConfig *refisconfig.Config
}
