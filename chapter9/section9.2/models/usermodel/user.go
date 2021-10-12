package usermodel

import (
	"database/sql"
	"errors"
	"time"

	"github.com/aliforever/golang-backend-training/chapter9/section9.2/srv/db"
)

type User struct {
	Id          int64
	Name        string
	Username    string
	Password    string
	WhenCreated time.Time
}

func FindById(id int) (data *User, err error) {
	if db.DB() == nil {
		err = errors.New("db_not_connected")
		return
	}

	var user = &User{}
	err = db.DB().QueryRow("select id, name, username, password, when_created from users WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.WhenCreated)
	if err != nil {
		return
	}
	data = user
	return
}

func FindAll() (users []User, err error) {
	if db.DB() == nil {
		err = errors.New("db_not_connected")
		return
	}

	var rows *sql.Rows

	rows, err = db.DB().Query("select id, name, username, password, when_created from users")
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
