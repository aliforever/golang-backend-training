package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section9.2/models"
	"github.com/aliforever/golang-backend-training/chapter9/section9.2/srv/db"
	"github.com/aliforever/golang-backend-training/chapter9/section9.2/srv/logger"

	_ "github.com/lib/pq"
)

func main() {
	log := logger.Begin()
	defer log.End()

	log.Log("Connecting to PostgreSQL")

	var err error

	err = db.Connect()
	if err != nil {
		log.Error("Error connecting to postgreSQL", err)
		return
	}
	defer db.Close()

	var id = 1
	log.Log("Finding user by id", id)

	var user models.User
	user, err = models.User{}.FindById(id)
	if err != nil {
		log.Error("Error find user by id", err)
		return
	}

	log.Trace("Found user", user)
}
