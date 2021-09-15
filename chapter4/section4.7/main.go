package main

import (
	"os"

	logger "github.com/aliforever/go-log"
)

var log *logger.Logger = logger.NewLogger(os.Stdout)

func main() {
	log.Begin()
	log.End()
}
