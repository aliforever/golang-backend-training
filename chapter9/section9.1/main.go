package main

import (
	"github.com/aliforever/golang-backend-training/chapter9/section9.1/models"
	"github.com/aliforever/golang-backend-training/chapter9/section9.1/srv/db"

	"github.com/aliforever/golang-backend-training/chapter9/section9.1/srv/logger"

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

	log.Log("Finding all users")

	var users []models.User
	users, err = models.User{}.FindAll()
	if err != nil {
		log.Error("Error finding users", err)
		return
	}
	log.Trace("List of users", users)
}
