package main

import (
	"net/http"
	"sync"

	"github.com/aliforever/golang-backend-training/chapter13/section13.1/handler"

	"github.com/aliforever/golang-backend-training/chapter13/section13.1/cfg"
	"github.com/aliforever/golang-backend-training/chapter13/section13.1/srv/logger"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go Case1()
	go Case2()
	wg.Wait()
}

func Case1() {
	defer wg.Done()
	log := logger.Begin()
	defer log.End()

	log.Inform("HTTP Address:", cfg.DefaultHTTPAddress)

	var h http.Handler

	h = handler.HelloWorldHandler{}

	err := http.ListenAndServe(cfg.DefaultHTTPAddress, h)
	if err != nil {
		log.ErrorF("Error listening on %s: %s", cfg.DefaultHTTPAddress, err)
		return
	}
}

func Case2() {
	defer wg.Done()
	log := logger.Begin()
	defer log.End()

	log.Inform("HTTP Address:", cfg.DefaultHTTPAddress2)

	var h http.Handler

	h = handler.HelloWorldHandler{}

	h = handler.LogHandler{
		SubHandler: h,
	}

	err := http.ListenAndServe(cfg.DefaultHTTPAddress2, h) // <----
	if nil != err {
		log.Error("Had problem with HTTP server:", err)
		return
	}
}
