package main

import (
	"net/http"
	"os"
	"strings"

	logger "github.com/aliforever/go-log"
)

var log *logger.Logger = logger.NewLogger(os.Stdout)

func main() {
	log.Log("App Started...")

	log.Inform("Registering routes")
	http.HandleFunc("/", ServeIndex)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("Error listening to port 8080", err)
		return
	}
	log.End()
}

func ServeIndex(writer http.ResponseWriter, request *http.Request) {
	log.Log("Received HTTP Request")
	log.Trace("HTTP Request", request)

	authorization := request.Header.Get("Authorization")
	if authorization == "" {
		log.Alert("User's authorization header is empty")
	} else {
		log.InformF("User's authorization header is %s", authorization)
	}

	userAgent := request.Header.Get("User-Agent")
	if strings.Contains(userAgent, "Googlebot") {
		log.HighlightF("Request is from a Googlebot: %s", userAgent)
	}

	writer.Write([]byte("Hello World"))
	log.Log("Written HTTP Response")
}
