package log

import (
	"context"
	"io"
	"sync"
	"sync/atomic"
)

var (
	loggerOnce  sync.Once
	loggerValue atomic.Value
)

type Logger interface {
	Writer() io.Writer
	With(...Field) Logger
	Debugf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Debug(ctx context.Context, message string)
	Error(ctx context.Context, message string)
	Panic(ctx context.Context, message string)
	Info(ctx context.Context, message string)
}

func getLogger() *logrusLogger {
	if loggerValue.Load() == nil {
		setLoggerOnce(defaultConfig)
	}
	return loggerValue.Load().(*logrusLogger)
}

func GetLogger() Logger {
	return getLogger()
}

func InitLogger(cfg config) {
	setLoggerOnce(cfg)
}
