package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section9.3/models/usermodel"
	"github.com/aliforever/golang-backend-training/chapter9/section9.3/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	var err error

	log.Log("Creating a new user")

	var id int
	id, err = usermodel.Create("Charles", "reiver2", "195487523")
	if err != nil {
		log.Error("Error creating a new user", err)
		return
	}

	log.Trace("Inserted User Id:", id)
}
