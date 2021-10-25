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
		fmt.Fprint(wr, "Hello World!")
		return
	case "/favorite/fruits":
		fmt.Fprint(wr, "apple banana cherry")
		return
	case "/about-us":
		fmt.Fprint(wr, "Lorem Ipsum")
		return
	case "/ip":
		fmt.Fprintf(wr, "Your IP is: %s", r.RemoteAddr)
		return
	case "/user-agent":
		fmt.Fprintf(wr, "Your User Agent is: %s", r.UserAgent())
		return
	case "/":
		fmt.Fprintf(wr, "Welcome!")
		return

	default:
		http.NotFound(wr, r)
		return
	}

}
