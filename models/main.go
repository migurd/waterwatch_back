package models

import (
	"database/sql"
	"time"
)

var db *sql.DB

// timeout of db transaction
const timeoutDB = time.Second * 3

type Models struct {
	// I wonder if this bullcrap will work
	Client Client
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}