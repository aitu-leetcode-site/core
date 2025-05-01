package log

import "context"

func With(fields ...Field) Logger {
	return getLogger().With(fields...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	getLogger().Debugf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	getLogger().Errorf(ctx, format, args...)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	getLogger().Panicf(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	getLogger().Infof(ctx, format, args...)
}

func Debug(ctx context.Context, message string) {
	getLogger().Debug(ctx, message)
}

func Error(ctx context.Context, message string) {
	getLogger().Error(ctx, message)
}

func Panic(ctx context.Context, message string) {
	getLogger().Panic(ctx, message)
}

func Info(ctx context.Context, message string) {
	getLogger().Info(ctx, message)
}
