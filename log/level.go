package log

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

const defaultLoggerLevel = logrus.InfoLevel

func getTrueCaller() string {
	const maxDepth = 20
	pcs := make([]uintptr, maxDepth)
	n := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.Function, "aitu-leetcode-site/core/log") {
			return frame.Function
		}
		if !more {
			break
		}
	}
	return "(unknown)"
}
