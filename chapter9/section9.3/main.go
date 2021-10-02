package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section9.3/models"
	"github.com/aliforever/golang-backend-training/chapter9/section9.3/srv/logger"

	_ "github.com/lib/pq"
)

func main() {
	log := logger.Begin()
	defer log.End()

	log.Log("Connecting to PostgreSQL")

	var err error

	log.Log("Creating a new user")

	var id int
	id, err = models.User{}.Create("Charles", "reiver2", "195487523")
	if err != nil {
		log.Error("Error creating a new user", err)
		return
	}

	log.Trace("Inserted User Id:", id)
}
