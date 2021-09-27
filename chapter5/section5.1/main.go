package main

import (
	"net/http"

	"github.com/aliforever/golang-backend-training/utils"

	"github.com/aliforever/golang-backend-training/chapter5/section5.1/srv/logger"
)

func main() {
	logger.DefaultLogger = logger.DefaultLogger.Begin()
	defer logger.DefaultLogger.End()
	handler := http.NewServeMux()
	handler.HandleFunc("/", IndexHandler)
	if err := http.ListenAndServe(":80", handler); err != nil {
		logger.DefaultLogger.Error("Error listening to port 80", err)
		return
	}
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	logger.DefaultLogger.Trace("Incoming HTTP Request", request)
	if err := utils.HttpOkRaw(writer, "Hello World!"); err != nil {
		logger.DefaultLogger.Error("Error writing response", err)
	}
}
