package db

import (
	"database/sql"
	"fmt"

	"github.com/aliforever/golang-backend-training/chapter9/section9.4/cfg"
	_ "github.com/lib/pq"
)

var databaseConnection *sql.DB

func DB() *sql.DB {
	return databaseConnection
}

func connect() (db *sql.DB, err error) {
	return sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName))
}

func init() {
	var err error
	databaseConnection, err = connect()
	if nil != err {
		panic(err)
	}
	if nil == databaseConnection {
		panic("could not connect to database")
	}
}
