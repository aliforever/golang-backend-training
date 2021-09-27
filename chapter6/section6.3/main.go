package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aliforever/golang-backend-training/chapter6/section6.3/srv/logger"

	"github.com/aliforever/golang-backend-training/utils"
)

func main() {
	logger.DefaultLogger = logger.DefaultLogger.Begin()
	defer logger.DefaultLogger.End()
	handler := http.NewServeMux()
	handler.HandleFunc("/", IndexHandler)
	handler.HandleFunc("/hello", HelloHandler)
	handler.HandleFunc("/addition", AdditionHandler)
	if err := http.ListenAndServe(":80", handler); err != nil {
		logger.DefaultLogger.Error("Error listening to port 80", err)
		return
	}
}

func AdditionHandler(writer http.ResponseWriter, request *http.Request) {
	logger.DefaultLogger.Trace("Incoming HTTP Request For Addition Handler", request)
	request.ParseForm()
	x, xErr := strconv.ParseInt(request.Form.Get("x"), 10, 64)
	if xErr != nil {
		logger.DefaultLogger.Trace("invalid_input_x", request)

		if err := utils.HttpBadRequestJSON(writer, "invalid_input_x"); err != nil {
			logger.DefaultLogger.Error("error writing response", err)
		}
		return
	}
	y, yErr := strconv.ParseInt(request.Form.Get("y"), 10, 64)
	if yErr != nil {
		logger.DefaultLogger.Trace("invalid_input_y", request)
		if err := utils.HttpBadRequestJSON(writer, "invalid_input_y"); err != nil {
			logger.DefaultLogger.Error("error writing response", err)
		}
		return
	}
	logger.DefaultLogger.LogF("Request's x & y were %d-%d", x, y)

	if err := utils.HttpOkJSON(writer, x+y); err != nil {
		logger.DefaultLogger.Error("error writing response", err)
	}
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	logger.DefaultLogger.Trace("Incoming HTTP Request For Name Handler", request)
	request.ParseForm()
	name := request.Form.Get("name")
	writer.Header().Set("Content-Type", "application/json")
	if name == "" {
		logger.DefaultLogger.Trace("Empty name passed", request)
		if err := utils.HttpBadRequestJSON(writer, "empty_name"); err != nil {
			logger.DefaultLogger.Error("error writing response", err)
		}
		return
	}
	logger.DefaultLogger.LogF("Request's name was %s", name)
	if err := utils.HttpOkJSON(writer, fmt.Sprintf("Hello %s!", name)); err != nil {
		logger.DefaultLogger.Error("error writing response", err)
	}
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	logger.DefaultLogger.Trace("Incoming HTTP Request For Index", request)
	if err := utils.HttpOkJSON(writer, "Hello World"); err != nil {
		logger.DefaultLogger.Error("error writing response", err)
	}
}
