package handler

import (
	"net/http"

	"github.com/aliforever/golang-backend-training/chapter13/section13.4/srv/logger"
)

type UserAgentLogger struct {
	SubHandler http.Handler
}

func (xhm UserAgentLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := logger.Begin()
	defer log.End()

	subHandler := xhm.SubHandler
	if nil == subHandler {
		log.Error("subhandler is not passed")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	userAgentHeader := r.Header.Get("User-Agent")
	if userAgentHeader == "" {
		return
	}

	log.Log("User-Agent For request is:", userAgentHeader)

	xhm.SubHandler.ServeHTTP(w, r)
}
