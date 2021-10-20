package handler

import (
	"net/http"

	"github.com/aliforever/golang-backend-training/chapter13/section13.1/srv/logger"
)

type LogHandler struct {
	SubHandler http.Handler
}

func (lh LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	subHandler := lh.SubHandler
	if nil == subHandler {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log := logger.Begin()
	subHandler.ServeHTTP(w, r)
	log.End()
}
