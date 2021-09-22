package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aliforever/go-log"
)

var logger = log.NewLogger(nil).Level(6)

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
		j, _ := json.Marshal(map[string]interface{}{
			"msg": "empty_name",
		})
		writer.Write(j)
		return
	}
	logger.LogF("Request's name was %s", name)
	j, _ := json.Marshal(map[string]interface{}{
		"msg": fmt.Sprintf("Hello %s!", name),
	})
	writer.Write(j)
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Trace("Incoming HTTP Request For Index", request)
	j, _ := json.Marshal(map[string]interface{}{
		"msg": "Hello World!",
	})

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(j)
}
