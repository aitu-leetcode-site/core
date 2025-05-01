package httpconfig

import (
	"github.com/aitu-leetcode-site/core/utils"
	"time"
)

type Config struct {
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Port         utils.Port
}
