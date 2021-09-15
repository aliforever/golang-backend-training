package main

import (
	"fmt"
	"os"

	"github.com/aliforever/golang-backend-training/chapter4/section4.2/libs/logger"
)

var log *logger.Logger = logger.NewLogger()

func main() {

	log.SetOutput(os.Stdout) // This is to set destination for log
	log.Begin()

	var name string = "Joe"
	log.LogF("The name was %q", name)

	fmt.Printf("Hello %s\n", name)
	log.End()

}
