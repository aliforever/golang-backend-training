package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aliforever/golang-backend-training/chapter13/section13.3/srv/logger"
)

var (
	internalServerError []byte = []byte("Internal Server Error") // <---------
)

type HelloWorldHandler struct{}

func (hwh HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := logger.Begin()
	defer log.End()

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Trace("received http request", r)
		log.Error("Error reading body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(internalServerError)
		return
	}

	log.InformF("Received Request Details:\n%s /\nHOST: %s\nX-HTTP-Method-Override: %s\nContent-Type: %s\nContent-Length: %s\n\n%s",
		r.Method, r.Host, r.Header.Get("X-HTTP-Method-Override"),
		r.Header.Get("Content-Type"), r.Header.Get("Content-Length"), b)
	fmt.Fprint(w, "Hello world!")
}
