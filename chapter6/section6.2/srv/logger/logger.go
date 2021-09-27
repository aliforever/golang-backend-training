package logger

import (
	"os"

	"github.com/aliforever/golang-backend-training/chapter6/section6.2/arg"

	logger "github.com/aliforever/go-log"
)

var DefaultLogger *logger.Logger = logger.NewLogger(os.Stdout)

func init() {
	DefaultLogger = DefaultLogger.Level(arg.LogLevel)
}
