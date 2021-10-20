package handler

import (
	"fmt"
	"net/http"

	"github.com/aliforever/golang-backend-training/chapter13/section13.4/srv/logger"
)

type HelloWorldHandler struct{}

func (hwh HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := logger.Begin()
	defer log.End()
	log.Log("running hello world handler")
	// log.Trace(r)
	fmt.Fprint(w, "Hello world!")
}
