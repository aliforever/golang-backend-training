package balancemodel

import (
	"errors"

	"github.com/aliforever/go-cad"

	"github.com/aliforever/golang-backend-training/chapter9/section10.8/srv/db"
)

type Balance struct {
	Id      int64
	Balance cad.CAD
}

func FindById(id int) (balance *Balance, err error) {
	if db.DB() == nil {
		err = errors.New("db_not_connected")
		return
	}

	balance = &Balance{}
	err = db.DB().QueryRow("select * from balance WHERE id=$1", id).Scan(&balance.Id, &balance.Balance)
	if err != nil {
		return
	}

	return
}

func Create(balance cad.CAD) (id int, err error) {
	if db.DB() == nil {
		err = errors.New("db_not_connected")
		return
	}

	query := `INSERT INTO balance (balance) VALUES ($1) RETURNING id`

	err = db.DB().QueryRow(query, balance).Scan(&id)
	if err != nil {
		return
	}

	return
}
