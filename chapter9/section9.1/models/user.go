package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/aliforever/golang-backend-training/chapter9/section9.1/srv/db"
)

type User struct {
	Id          int64
	Name        string
	Username    string
	Password    string
	WhenCreated time.Time
}

func (User) FindAll() (users []User, err error) {
	if db.DB() == nil {
		err = errors.New("db_not_connected")
		return
	}

	var rows *sql.Rows

	rows, err = db.DB().Query("select * from users")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user = User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.WhenCreated)
		if err != nil {
			return
		}
		users = append(users, user)
	}

	err = rows.Err()

	return
}
