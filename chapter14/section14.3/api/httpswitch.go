package api

import (
	"fmt"
	"net/http"
)

type HTTPSwitch struct {
}

func (hs HTTPSwitch) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/hello":
		fmt.Fprintf(wr, "Hello world! method = %q", r.Method)
		return
	default:
		http.NotFound(wr, r)
		return
	}

}
