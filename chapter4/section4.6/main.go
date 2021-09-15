package main

import (
	"os"

	logger "github.com/aliforever/go-log"
)

var log *logger.Logger = logger.NewLogger()

func main() {
	log.SetOutput(os.Stdout) // This is to set destination for log
	log.Begin()
	log.End()
}
