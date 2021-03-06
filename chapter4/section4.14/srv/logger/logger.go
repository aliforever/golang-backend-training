package logger

import (
	"os"

	"github.com/aliforever/golang-backend-training/chapter4/section4.14/arg"

	logger "github.com/aliforever/go-log"
)

var DefaultLogger *logger.Logger = logger.NewLogger(os.Stdout)

func init() {
	DefaultLogger = DefaultLogger.Level(arg.LogLevel)
}
