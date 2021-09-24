package main

import (
	"fmt"
	"net/http"

	"github.com/aliforever/golang-backend-training/utils"

	log "github.com/aliforever/go-log"
)

var logger = log.NewLogger(nil).Level(utils.LogLevel)

func main() {
	logger = logger.Begin()
	defer logger.End()
	handler := http.NewServeMux()
	handler.HandleFunc("/", IndexHandler)
	handler.HandleFunc("/hello", NameHandler)
	if err := http.ListenAndServe(":80", handler); err != nil {
		logger.Error("Error listening to port 80", err)
		return
	}
}

func NameHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Trace("Incoming HTTP Request For Name Handler", request)
	request.ParseForm()
	name := request.Form.Get("name")
	if name == "" {
		logger.Trace("Empty name passed", request)
		writer.Write([]byte("empty_name"))
		return
	}
	logger.LogF("Request's name was %s", name)
	writer.Write([]byte(fmt.Sprintf("Hello %s!", name)))
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Trace("Incoming HTTP Request For Index", request)
	writer.Write([]byte("Hello World"))
}
