package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section9.5/models/usermodel"
	"github.com/aliforever/golang-backend-training/chapter9/section9.5/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	log.Log("Connecting to PostgreSQL")

	var err error

	log.Log("Deleting user by id")
	var id = 5

	err = usermodel.DeleteById(id)
	if err != nil {
		log.Error("Error deleting user", err)
		return
	}
	log.Log("Successfully deleted user")
}
