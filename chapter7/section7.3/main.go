package main

import (
	"encoding/json"
	"net/http"

	"github.com/aliforever/golang-backend-training/chapter7/section7.3/srv/logger"
)

type Response struct {
	StatusCode   int64  `json:"status_code"`
	StatusName   string `json:"status_name"`
	HumanMessage string `json:"message"`
}

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
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(Response{
		StatusCode:   http.StatusNotFound,
		StatusName:   "Not Found",
		HumanMessage: "This is not the URI you are looking for.",
	})
	if err != nil {
		log.Error("Error encoding json response", err)
		writer.Header().Set("Content-Type", "application/text")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
}
