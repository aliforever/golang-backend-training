package main

import (
	"net/http"

	"github.com/aliforever/golang-backend-training/chapter6/section6.1/srv/logger"

	"github.com/aliforever/golang-backend-training/utils"
)

func main() {
	log := logger.Begin()
	defer log.End()
	handler := http.NewServeMux()
	handler.HandleFunc("/", IndexHandler)
	if err := http.ListenAndServe(":80", handler); err != nil {
		log.Error("Error listening to port 80", err)
		return
	}
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	log := logger.Begin()
	defer log.End()
	log.Trace("Incoming HTTP Request For Index", request)

	if err := utils.HttpOkJSON(writer, "Hello World!"); err != nil {
		log.Error("error writing response", err)
	}
}
