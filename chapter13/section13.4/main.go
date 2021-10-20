package main

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aliforever/golang-backend-training/chapter13/section13.4/api"
	"github.com/aliforever/golang-backend-training/chapter13/section13.4/cfg"

	"github.com/aliforever/golang-backend-training/chapter13/section13.4/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	log.Log("App Started")

	go testServer()

	err := http.ListenAndServe(cfg.DefaultHTTPAddress, api.UserAgentLogger{SubHandler: api.HelloWorldHandler{}})

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

	req.Header.Set("User-Agent", "CustomUserAgent/2045")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.ErrorF("can't post to %s: %s", cfg.DefaultHTTPAddress, err)
		return
	}

	log.InformF("Sent Request User Agent: %s", resp.Request.Header.Get("User-Agent"))
}
