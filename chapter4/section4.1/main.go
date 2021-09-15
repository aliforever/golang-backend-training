package main

import (
	"fmt"
	"os"

	"github.com/aliforever/golang-backend-training/chapter4/section4.1/libs/logger"
)

var log *logger.Logger = logger.NewLogger()

func main() {
	log.SetOutput(os.Stdout) // This is to set destination for log
	log.Begin()

	var msg string = "Hello World!"
	fmt.Println(msg)

	log.Log("I said:", msg)
	log.End()

}
