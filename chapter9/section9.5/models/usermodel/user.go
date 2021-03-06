package usermodel

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aliforever/golang-backend-training/chapter9/section9.5/srv/db"
)

type User struct {
	Id          int64
	Name        string
	Username    string
	Password    string
	WhenCreated time.Time
}

func DeleteById(id int) (err error) {
	query := `DELETE FROM users WHERE id = $1`
	_, err = db.DB().Exec(query, id)
	return
}

func updateByIdQuery(query *strings.Builder, id int, data map[string]interface{}) (values []interface{}, err error) {
	if 0 >= len(data) {
		err = errors.New("bad_request")
		return
	}

	values = make([]interface{}, len(data)+1)
	values[0] = id

	query.WriteString(`UPDATE users `)

	counter := 2
	separator := "SET "
	for key, value := range data {

		query.WriteString(separator)
		separator = ", "

		query.WriteString(key)
		query.WriteString(" = $")
		fmt.Fprintf(query, "%d", counter)

		values[counter-1] = value
		counter++
	}
	query.WriteString(" WHERE id=$1")
	return
}

func UpdateById(id int, data map[string]interface{}) (err error) {
	var (
		query  strings.Builder
		values []interface{}
	)

	values, err = updateByIdQuery(&query, id, data)
	if err != nil {
		return
	}

	_, err = db.DB().Exec(query.String(), values...)
	return
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
