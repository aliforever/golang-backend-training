package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aliforever/golang-backend-training/chapter14/section14.3/api"
	"github.com/aliforever/golang-backend-training/chapter14/section14.3/cfg"
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

	methods := []string{
		"GET", "PATCH", "POST", "DELETE", "CUSTOM",
	}

	for _, method := range methods {
		err := TestRoute(method)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error testing route %s", err))
			continue
		}
	}
}

func TestRoute(method string) (err error) {
	var req *http.Request
	req, err = http.NewRequest(method, "http://localhost"+cfg.DefaultAddress+"/hello", nil)
	if err != nil {
		return
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	expected := fmt.Sprintf("Hello world! method = %q", method)
	if string(b) != expected {
		err = errors.New(fmt.Sprintf("Invalid response.\nExpected: %s\nActual: %s", expected, b))
	}
	return
}
