package models

import (
	"context"
)

type Account struct {
	ClientID int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

func (a *Account) CreateAccount() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO account 
		(client_id, username, password, status)
		VALUES ($1, $2, $3, $4)`

	_, err := db.QueryContext(
		ctx,
		query,
		a.ClientID,
		a.Username,
		a.Password,
		a.Status,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) DoesAccountExist() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT id FROM account WHERE id = ?`

	var id int64
	row := db.QueryRowContext(ctx, query, a.ClientID)
	err := row.Scan(&id)
	if err != nil { // if it doesn't return any row, then error
		return false, nil
	}
	return true, nil
}
