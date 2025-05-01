package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

var loggerKeys = []string{
	"level",
	"timestamp",
	"caller",
}

var internalKeysGetters = map[string]func() interface{}{
	"level": func() interface{} {
		return logrus.InfoLevel
	},
	"timestamp": func() interface{} {
		return time.Now()
	},
	"caller": func() interface{} {
		return getTrueCaller()
	},
}

func ctxToLoggerKeys(ctx context.Context) logrus.Fields {
	fields := make(logrus.Fields)
	for _, internalKey := range loggerKeys {
		ctxValue := ctx.Value(internalKey)
		if ctxValue == nil {
			ctxValue = internalKeysGetters[internalKey]()
		}
		fields[internalKey] = ctxValue
	}
	return fields
}
