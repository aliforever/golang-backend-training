package handler

import (
	"net/http"

	"github.com/aliforever/golang-backend-training/chapter13/section13.3/srv/logger"
)

type XHTTPMethodOverrideHandler struct {
	SubHandler http.Handler
}

func (xhm XHTTPMethodOverrideHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := logger.Begin()
	defer log.End()

	subHandler := xhm.SubHandler
	if nil == subHandler {
		log.Error("subhandler is not passed")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	xHttpHeader := r.Header.Get("X-HTTP-Method-Override")
	if xHttpHeader == "" {
		return
	}

	r.Method = xHttpHeader
	r.Header.Del("X-HTTP-Method-Override")

	xhm.SubHandler.ServeHTTP(w, r)
}
