package logger

import (
	"os"

	"github.com/aliforever/golang-backend-training/chapter13/section13.3/arg"

	logger "github.com/aliforever/go-log"
)

var defaultLogger *logger.Logger = logger.NewLogger(os.Stdout)

func init() {
	defaultLogger = defaultLogger.Level(arg.LogLevel)
}

func Begin() *logger.Logger {
	return defaultLogger.Begin()
}

type ForMiddleware struct {
}

func (fm ForMiddleware) Error(data ...interface{}) {
	l := Begin()
	defer l.End()

	l.Error(data...)
}
