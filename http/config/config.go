package httpconfig

import (
	"github.com/aitu-leetcode-site/core/utils/port"
	"time"
)

type Config struct {
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	Port         port.Port
}
