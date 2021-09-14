package main

import (
	"github.com/aliforever/golang-backend-training/chapter1/section1.9/cfg"
	"github.com/aliforever/golang-backend-training/chapter1/section1.9/lib/formletter"
)

func main() {
	formletter.FormLetter(cfg.Name, cfg.Weather, cfg.Snack)
}
