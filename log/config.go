package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type config struct {
	writer      io.Writer
	loggerLevel logrus.Level
}

var defaultConfig = config{
	writer:      os.Stdout,
	loggerLevel: defaultLoggerLevel,
}
