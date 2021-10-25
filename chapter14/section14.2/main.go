package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aliforever/golang-backend-training/chapter14/section14.2/api"
	"github.com/aliforever/golang-backend-training/chapter14/section14.2/cfg"
)

func main() {
	go func() {
		err := http.ListenAndServe(cfg.DefaultAddress, api.HTTPSwitch{})
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	time.Sleep(time.Second * 2)
	routes := map[string]string{
		"/hello":           "Hello World!",
		"/favorite/fruits": "apple banana cherry",
		"/about-us":        "Lorem Ipsum",
		"/":                "Welcome!",
		"/test_not_found":  "404 page not found\n",
	}

	for route, response := range routes {
		err := TestRoute(route, response)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error testing route %s : %s", route, err))
			continue
		}
	}
}

func TestRoute(path string, response string) (err error) {
	var resp *http.Response
	resp, err = http.Get("http://localhost" + cfg.DefaultAddress + path)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if string(b) != response {
		err = errors.New(fmt.Sprintf("Invalid response.\nExpected: %s\nActual: %s", response, b))
	}
	return
}
