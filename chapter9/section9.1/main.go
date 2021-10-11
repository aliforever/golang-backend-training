package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section9.1/models/usermodel"
	"github.com/aliforever/golang-backend-training/chapter9/section9.1/srv/logger"
)

func main() {
	log := logger.Begin()
	defer log.End()

	var err error

	log.Log("Finding all users")
	var users []usermodel.User
	users, err = usermodel.FindAll()
	if err != nil {
		log.Error("Error finding users", err)
		return
	}
	log.Trace("List of users", users)
}
