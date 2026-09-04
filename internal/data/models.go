package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")
var ErrEditConflict = errors.New("edit conflict")

type Model struct {
	Movies MovieModel
	Users  UserModel
	Tokens TokenModel
}

func NewModel(db *sql.DB) Model {
	return Model{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db},
		Tokens: TokenModel{DB: db},
	}
}
