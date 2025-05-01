package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
)

type logrusLogger struct {
	entry *logrus.Entry
}

func newLogrusLogger(cfg config) *logrusLogger {
	instance := logrus.New()
	instance.Level = cfg.loggerLevel
	instance.Out = cfg.writer
	return &logrusLogger{
		entry: instance.WithContext(context.Background()),
	}
}

func (l *logrusLogger) getLoggerWithCtx(ctx context.Context) *logrusLogger {
	ctxFields := ctxToLoggerKeys(ctx)
	loggerEntry := l.entry.WithFields(ctxFields).
		WithContext(ctx)
	return &logrusLogger{entry: loggerEntry}
}

func setLoggerOnce(cfg config) {
	loggerOnce.Do(func() {
		logger := newLogrusLogger(cfg)
		loggerValue.Store(logger)
	})
}

func (l *logrusLogger) Writer() io.Writer {
	return l.entry.Logger.Out
}
func (l *logrusLogger) With(fields ...Field) Logger {
	logrusEntry := l.entry.WithFields(
		convertFields(fields),
	)
	return &logrusLogger{entry: logrusEntry}
}

func (l *logrusLogger) Printf(ctx context.Context, format string, args ...interface{}) {
	l.getLoggerWithCtx(ctx).entry.Printf(format, args...)
}

func (l *logrusLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.getLoggerWithCtx(ctx).entry.Debugf(format, args...)
}

func (l *logrusLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.getLoggerWithCtx(ctx).entry.Errorf(format, args...)
}

func (l *logrusLogger) Panicf(ctx context.Context, format string, args ...interface{}) {
	l.getLoggerWithCtx(ctx).entry.Panicf(format, args...)
}

func (l *logrusLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.getLoggerWithCtx(ctx).entry.Infof(format, args...)
}

func (l *logrusLogger) Debug(ctx context.Context, message string) {
	l.getLoggerWithCtx(ctx).entry.Debug(message)
}

func (l *logrusLogger) Error(ctx context.Context, message string) {
	l.getLoggerWithCtx(ctx).entry.Error(message)
}

func (l *logrusLogger) Panic(ctx context.Context, message string) {
	l.getLoggerWithCtx(ctx).entry.Panic(message)
}

func (l *logrusLogger) Info(ctx context.Context, message string) {
	l.getLoggerWithCtx(ctx).entry.Info(message)
}
