package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section9.2/models/usermodel"
	"github.com/aliforever/golang-backend-training/chapter9/section9.2/srv/logger"

	_ "github.com/lib/pq"
)

func main() {
	log := logger.Begin()
	defer log.End()

	log.Log("Connecting to PostgreSQL")

	var err error

	var id = 1
	log.Log("Finding user by id", id)

	var user usermodel.User
	user, err = usermodel.FindById(id)
	if err != nil {
		log.Error("Error find user by id", err)
		return
	}

	log.Trace("Found user", user)
}
