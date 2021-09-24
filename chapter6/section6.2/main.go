package main

import (
	"fmt"
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
	handler.HandleFunc("/hello", HelloHandler)
	if err := http.ListenAndServe(":80", handler); err != nil {
		logger.Error("Error listening to port 80", err)
		return
	}
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Trace("Incoming HTTP Request For Name Handler", request)
	request.ParseForm()
	name := request.Form.Get("name")
	writer.Header().Set("Content-Type", "application/json")
	if name == "" {
		logger.Trace("Empty name passed", request)

		if err := utils.HttpBadRequestJSON(writer, "empty_name"); err != nil {
			logger.Error("error writing response", err)
		}
		return
	}
	utils.HttpOkJSON(writer, fmt.Sprintf("Hello %s!", name))
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Trace("Incoming HTTP Request For Index", request)

	if err := utils.HttpOkJSON(writer, "Hello World!"); err != nil {
		logger.Error("error writing response", err)
	}
}
