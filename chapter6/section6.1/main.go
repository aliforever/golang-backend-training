package main

import (
	"net/http"

	"github.com/aliforever/golang-backend-training/utils"

	"github.com/aliforever/go-log"
)

var logger = log.NewLogger(nil).Level(utils.LogLevel)

func main() {
	logger = logger.Begin()
	defer logger.End()
	handler := http.NewServeMux()
	handler.HandleFunc("/", IndexHandler)
	if err := http.ListenAndServe(":80", handler); err != nil {
		logger.Error("Error listening to port 80", err)
		return
	}
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Trace("Incoming HTTP Request For Index", request)

	if err := utils.HttpOkJSON(writer, "Hello World!"); err != nil {
		logger.Error("error writing response", err)
	}
}
