package redis

import "github.com/redis/go-redis/v9"

type (
	IntCmd             = redis.IntCmd
	FloatCmd           = redis.FloatCmd
	StatusCmd          = redis.StatusCmd
	StringCmd          = redis.StringCmd
	DurationCmd        = redis.DurationCmd
	BoolCmd            = redis.BoolCmd
	SliceCmd           = redis.SliceCmd
	StringSliceCmd     = redis.StringSliceCmd
	MapStringStringCmd = redis.MapStringStringCmd
)
