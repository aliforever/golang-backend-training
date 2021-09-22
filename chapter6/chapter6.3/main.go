package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aliforever/go-log"
)

var logger = log.NewLogger(nil).Level(6)

func main() {
	logger = logger.Begin()
	defer logger.End()
	handler := http.NewServeMux()
	handler.HandleFunc("/", IndexHandler)
	handler.HandleFunc("/hello", HelloHandler)
	handler.HandleFunc("/addition", AdditionHandler)
	if err := http.ListenAndServe(":80", handler); err != nil {
		logger.Error("Error listening to port 80", err)
		return
	}
}

func writeResponse(writer http.ResponseWriter, message interface{}) {
	j, _ := json.Marshal(map[string]interface{}{
		"msg": message,
	})
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(j)
	return
}

func AdditionHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Trace("Incoming HTTP Request For Addition Handler", request)
	request.ParseForm()
	x, xErr := strconv.ParseInt(request.Form.Get("x"), 10, 64)
	if xErr != nil {
		logger.Trace("invalid_input_x", request)
		writeResponse(writer, "invalid_input_x")
		return
	}
	y, yErr := strconv.ParseInt(request.Form.Get("y"), 10, 64)
	if yErr != nil {
		logger.Trace("invalid_input_y", request)
		writeResponse(writer, "invalid_input_y")
		return
	}
	logger.LogF("Request's x & y were %d-%d", x, y)
	writeResponse(writer, x+y)
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Trace("Incoming HTTP Request For Name Handler", request)
	request.ParseForm()
	name := request.Form.Get("name")
	writer.Header().Set("Content-Type", "application/json")
	if name == "" {
		logger.Trace("Empty name passed", request)
		writeResponse(writer, "empty_name")
		return
	}
	logger.LogF("Request's name was %s", name)
	writeResponse(writer, fmt.Sprintf("Hello %s!", name))
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Trace("Incoming HTTP Request For Index", request)
	writeResponse(writer, "Hello World")
}
