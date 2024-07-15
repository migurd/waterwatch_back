package controllers

import (
	"database/sql"
)

var db *sql.DB

type Controllers struct {
}

func New(dbConn *sql.DB) Controllers {
	db = dbConn
	return Controllers{}
}
