package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section9.1/models"
	"github.com/aliforever/golang-backend-training/chapter9/section9.1/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	var err error

	log.Log("Finding all users")
	var users []models.User
	users, err = models.User{}.FindAll()
	if err != nil {
		log.Error("Error finding users", err)
		return
	}
	log.Trace("List of users", users)
}
