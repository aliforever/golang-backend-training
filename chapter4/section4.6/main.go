package main

import (
	"flag"

	logger "github.com/aliforever/go-log"
)

var log *logger.Logger = logger.NewLogger(nil)

func processFlags(verbose *bool) {
	flag.BoolVar(verbose, "v", false, "--v")
}

func main() {
	var v bool
	processFlags(&v)

	if v {
		log.Begin()
		log.End()
	}
}
