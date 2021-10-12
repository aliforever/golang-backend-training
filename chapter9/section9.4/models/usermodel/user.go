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
	if 0 >= len(data) {
		err = errors.New("bad_request")
		return
	}

	var values []interface{} = make([]interface{}, len(data)+1)
	values[0] = id

	var query strings.Builder
	query.WriteString(`UPDATE users SET `)

	counter := 2
	lengthCounter := 0
	separator := ", "
	for key, value := range data {
		if lengthCounter == len(data)-1 {
			separator = ""
		}
		fmt.Fprintf(&query, "%s = $%d%s", key, counter, separator)
		lengthCounter++
		values[counter-1] = value
		counter++
	}
	query.WriteString(" WHERE id = $1")
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
