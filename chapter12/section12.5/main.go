package main

import "github.com/aliforever/golang-backend-training/chapter12/section12.5/srv/logger"

func main() {
	log := logger.Begin()
	defer log.End()

	log.Inform("App Started")
}
