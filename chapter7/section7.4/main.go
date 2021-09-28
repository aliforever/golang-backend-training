package main

import (
	"net/http"

	"github.com/aliforever/go-httpjson"

	"github.com/aliforever/golang-backend-training/chapter7/section7.4/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	handler := http.NewServeMux()
	handler.HandleFunc("/", NotFoundHandler)

	err := http.ListenAndServe(":80", handler)
	if err != nil {
		log.Error("error listening on port 80", err)
		return
	}
}

func NotFoundHandler(writer http.ResponseWriter, _ *http.Request) {
	log := logger.Begin()
	defer log.End()

	err := httpjson.NotFound(writer, "This is not the URI you are looking for.")
	if err != nil {
		log.Error("Error serving response", err)
	}
}
