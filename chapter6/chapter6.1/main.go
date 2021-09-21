package main

import (
	"encoding/json"
	"net/http"

	"github.com/aliforever/go-log"
)

var logger = log.NewLogger(nil).Level(6)

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
	j, _ := json.Marshal(map[string]interface{}{
		"msg": "Hello World!",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(j)
}
