package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section9.4/models/usermodel"
	"github.com/aliforever/golang-backend-training/chapter9/section9.4/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	log.Log("Connecting to PostgreSQL")

	var err error

	log.Log("Updating user by id")
	var id = 5

	err = usermodel.UpdateById(id, map[string]interface{}{
		"username": "new_username",
	})
	if err != nil {
		log.Error("Error updating user", err)
		return
	}
	log.Log("Successfully updated user")
}
