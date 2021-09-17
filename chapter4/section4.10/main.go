package main

import (
	"os"
	"time"

	logger "github.com/aliforever/go-log"
)

var log *logger.Logger = logger.NewLogger(os.Stdout)

func main() {
	localLogger := log.Begin()
	localLogger.Log("Hello World")
	time.Sleep(time.Millisecond * 4)
	localLogger.End()
}
