package main

import "github.com/aliforever/golang-backend-training/chapter4/section4.14/srv/logger"

func main() {
	logger.DefaultLogger.Alert("This is an Alert")
	logger.DefaultLogger.Error("This is an Error")
	logger.DefaultLogger.Warn("This is a Warning")
	logger.DefaultLogger.Highlight("This is a highlight")
	logger.DefaultLogger.InformF("This is an information")
	logger.DefaultLogger.Log("This is a log")
	logger.DefaultLogger.Trace("This is a trace")
}
