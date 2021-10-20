package main

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aliforever/go-xhttpmiddleware"

	"github.com/aliforever/golang-backend-training/chapter13/section13.3/api"
	"github.com/aliforever/golang-backend-training/chapter13/section13.3/cfg"

	"github.com/aliforever/golang-backend-training/chapter13/section13.3/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	log.Log("App Started")

	go testServer()

	err := http.ListenAndServe(cfg.DefaultHTTPAddress, xhttpmiddleware.NewXHTTPMethodOverrideHandler(api.HelloWorldHandler{}, logger.ForMiddleware{}))

	if err != nil {
		log.ErrorF("Error listening to %s: %s", cfg.DefaultHTTPAddress, err)
		return
	}
}

func testServer() {
	log := logger.Begin()
	defer log.End()

	time.Sleep(time.Second * 2)

	data := url.Values{}
	data.Set("phone_number", "604-555-5555")

	req, err := http.NewRequest("POST", "http://localhost"+cfg.DefaultHTTPAddress, strings.NewReader(data.Encode()))
	if err != nil {
		log.ErrorF("can't create request to %s: %s", cfg.DefaultHTTPAddress, err)
		return
	}

	req.Header.Set("X-HTTP-Method-Override", "PATCH")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.ErrorF("can't post to %s: %s", cfg.DefaultHTTPAddress, err)
		return
	}

	log.InformF("Sent Request Details:\n%s /\nHOST: %s\nX-HTTP-Method-Override: %s\nContent-Type: %s\nContent-Length: %s\n\n%s",
		resp.Request.Method, resp.Request.Host, resp.Request.Header.Get("X-HTTP-Method-Override"),
		resp.Request.Header.Get("Content-Type"), resp.Header.Get("Content-Length"), data.Encode())
}
