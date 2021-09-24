package main

import (
	"database/sql"
	"fmt"

	"github.com/aliforever/golang-backend-training/utils"

	"github.com/aliforever/go-log"

	"github.com/aliforever/golang-backend-training/chapter9/section9.1/models"

	_ "github.com/lib/pq"
)

var logger = log.NewLogger(nil).Level(utils.LogLevel)

func main() {
	var (
		dbHost     = "localhost"
		dbPort     = 5432
		dbName     = "my_database"
		dbUser     = "postgres"
		dbPassword = "159852"
	)

	logger.Log("Connecting to PostgreSQL")
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName))
	if err != nil {
		logger.Error("Error connecting to PostgreSQL", err)
		return
	}
	defer db.Close()

	logger.Log("Finding all users")
	rows, err := db.Query("select * from users")
	if err != nil {
		logger.Error("Error finding all users", err)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user = models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.WhenCreated)
		if err != nil {
			logger.Log("Error scanning user", err)
			return
		}
		users = append(users, user)
	}

	logger.Trace("List of users", users)
}
