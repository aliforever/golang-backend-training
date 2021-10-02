package log

import (
	"github.com/aliforever/go-log"
	"github.com/aliforever/golang-backend-training/chapter10/section10.3/arg"
)

var defaultLogger *log.Logger

func init() {
	defaultLogger = log.NewLogger(nil).Level(arg.LogLevel)
}

func Begin() *log.Logger {
	return defaultLogger.Begin()
}
