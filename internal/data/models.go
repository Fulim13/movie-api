package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")
var ErrEditConflict = errors.New("edit conflict")

type Model struct {
	Movies MovieModel
}

func NewModel(db *sql.DB) Model {
	return Model{
		Movies: MovieModel{DB: db},
	}
}
