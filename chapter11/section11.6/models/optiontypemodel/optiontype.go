package optiontypemodel

import (
	"errors"

	"github.com/aliforever/golang-backend-training/chapter11/section11.6/lib/int64option"

	"github.com/aliforever/golang-backend-training/chapter11/section11.6/srv/db"
)

type OptionType struct {
	Id         int64
	Name       string
	OptionType int64option.Type
}

func Create(name string, optionType int64option.Type) (id int, err error) {
	if db.DB() == nil {
		err = errors.New("db_not_connected")
		return
	}

	query := `INSERT INTO option_types (name, option_type) VALUES ($1, $2) RETURNING id`

	err = db.DB().QueryRow(query, name, optionType).Scan(&id)
	if err != nil {
		return
	}

	return
}
