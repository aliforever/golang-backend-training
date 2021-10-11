package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section10.9/models/balancemodel"
	"github.com/aliforever/golang-backend-training/chapter9/section10.9/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	log.Log("Connecting to PostgreSQL")

	var err error

	log.Log("Finding balance by id")
	var id = 1

	var balance *balancemodel.Balance

	balance, err = balancemodel.FindById(id)
	if err != nil {
		log.Error("Error finding balance", err)
		return
	}

	log.Trace("Balance:", balance)
}
