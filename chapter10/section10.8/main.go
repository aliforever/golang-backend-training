package main

import (
	"github.com/aliforever/go-cad"
	"github.com/aliforever/golang-backend-training/chapter9/section10.8/models/balancemodel"
	"github.com/aliforever/golang-backend-training/chapter9/section10.8/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	var err error

	log.Log("Creating a new balance")

	balance := cad.Cents(1026)

	var id int
	id, err = balancemodel.Create(balance)
	if err != nil {
		log.Error("Error creating a new balance", err)
		return
	}

	log.Trace("Inserted balance Id:", id)
}
