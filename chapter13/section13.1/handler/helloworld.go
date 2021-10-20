package handler

import (
	"fmt"
	"net/http"
)

type HelloWorldHandler struct{}

func (hwh HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
