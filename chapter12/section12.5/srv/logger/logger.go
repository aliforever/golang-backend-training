package logger

import (
	"os"

	"github.com/aliforever/golang-backend-training/chapter12/section12.5/arg"

	logger "github.com/aliforever/go-log"
)

var defaultLogger *logger.Logger = logger.NewLogger(os.Stdout)

func init() {
	defaultLogger = defaultLogger.Level(arg.LogLevel)
}

func Begin() *logger.Logger {
	return defaultLogger.Begin()
}