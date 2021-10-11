package usermodel

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aliforever/golang-backend-training/chapter9/section9.4/srv/db"
)

type User struct {
	Id          int64
	Name        string
	Username    string
	Password    string
	WhenCreated time.Time
}

func UpdateById(id int, data map[string]interface{}) (err error) {
	query := `UPDATE users SET `

	var (
		queryParts []string
		values     []interface{}
	)
	counter := 2
	for key, value := range data {
		queryParts = append(queryParts, fmt.Sprintf("%s = $%d", key, counter))
		values = append(values, value)
		counter++
	}
	values = append([]interface{}{id}, values...)

	query += strings.Join(queryParts, ",") + " WHERE id = $1"
	_, err = db.DB().Exec(query, values...)
	return
}

func FindById(id int) (user *User, err error) {
	if db.DB() == nil {
		err = errors.New("db_not_connected")
		return
	}

	user = &User{}
	err = db.DB().QueryRow("select * from users WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.WhenCreated)
	if err != nil {
		return
	}

	return
}

func Create(name, username, password string) (id int, err error) {
	if db.DB() == nil {
		err = errors.New("db_not_connected")
		return
	}

	query := `INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id`

	err = db.DB().QueryRow(query, name, username, password).Scan(&id)
	if err != nil {
		return
	}

	return
}

func FindAll() (users []User, err error) {
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
