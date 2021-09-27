package main

import (
	"net/http"

	"github.com/aliforever/golang-backend-training/chapter7/section7.4/srv/logger"
)

type Response struct {
	StatusCode   int64  `json:"status_code"`
	StatusName   string `json:"status_name"`
	HumanMessage string `json:"message"`
}

func main() {
	logger.DefaultLogger = logger.DefaultLogger.Begin()
	defer logger.DefaultLogger.End()

	handler := http.NewServeMux()
	handler.HandleFunc("/", NotFoundHandler)

	err := http.ListenAndServe(":80", handler)
	if err != nil {
		logger.DefaultLogger.Error("error listening on port 80", err)
		return
	}
}

func NotFoundHandler(writer http.ResponseWriter, _ *http.Request) {
	// TODO
}
