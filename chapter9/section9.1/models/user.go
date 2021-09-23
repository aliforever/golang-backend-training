package models

import "time"

type User struct {
	Id          int64
	Name        string
	Username    string
	Password    string
	WhenCreated time.Time
}
