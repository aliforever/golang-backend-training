package main

import (
	"net/http"

	"github.com/aliforever/golang-backend-training/chapter5/section5.2/srv/logger"

	"github.com/aliforever/golang-backend-training/utils"
)

func main() {
	logger.DefaultLogger = logger.DefaultLogger.Begin()
	defer logger.DefaultLogger.End()
	handler := http.NewServeMux()
	handler.HandleFunc("/", IndexHandler)
	handler.HandleFunc("/hello", NameHandler)
	if err := http.ListenAndServe(":80", handler); err != nil {
		logger.DefaultLogger.Error("Error listening to port 80", err)
		return
	}
}

func NameHandler(writer http.ResponseWriter, request *http.Request) {
	logger.DefaultLogger.Trace("Incoming HTTP Request For Name Handler", request)
	request.ParseForm()
	name := request.Form.Get("name")
	if name == "" {
		logger.DefaultLogger.Trace("Empty name passed", request)
		if err := utils.HttpBadRequestRaw(writer, "empty_name"); err != nil {
			logger.DefaultLogger.Error("Error writing response", err)
		}
		return
	}
	logger.DefaultLogger.LogF("Request's name was %s", name)
	if err := utils.HttpOkRaw(writer, "empty_name"); err != nil {
		logger.DefaultLogger.Error("Error writing response", err)
	}
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	logger.DefaultLogger.Trace("Incoming HTTP Request For Index", request)
	if err := utils.HttpOkRaw(writer, "Hello World!"); err != nil {
		logger.DefaultLogger.Error("Error writing response", err)
	}
}
