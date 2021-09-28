package db

import (
	"database/sql"
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter9/section9.2/cfg"
)

var databaseConnection *sql.DB

func DB() *sql.DB {
	return databaseConnection
}

func Connect() (err error) {
	databaseConnection, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName))
	return err
}

func Close() {
	databaseConnection.Close()
}
