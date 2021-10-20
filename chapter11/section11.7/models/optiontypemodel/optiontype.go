package optiontypemodel

import (
	"errors"

	"github.com/aliforever/golang-backend-training/chapter11/section11.7/lib/int64option"

	"github.com/aliforever/golang-backend-training/chapter11/section11.7/srv/db"
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
