package optiontypemodel

import (
	"errors"

	"github.com/aliforever/golang-backend-training/chapter11/section11.6/models/int64option"

	"github.com/aliforever/golang-backend-training/chapter11/section11.6/srv/db"
)

type OptionType struct {
	Id         int64
	Name       string
	OptionType int64option.Type
}

// CREATE TABLE option_types (id BIGSERIAL PRIMARY KEY,name VARCHAR(50) NOT NULL,option_type BIGINT UNIQUE NOT NULL);

func FindById(id int) (optionType *OptionType, err error) {
	if db.DB() == nil {
		err = errors.New("db_not_connected")
		return
	}

	var user = &OptionType{}
	err = db.DB().QueryRow("select id, name, option_type from option_types WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.OptionType)
	if err == nil {
		optionType = user
	}
	return
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
