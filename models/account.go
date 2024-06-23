package models

import "context"

type Account struct {
	UserID   int64  `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

func (a *Account) CreateAccount() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO account 
		(user_id, username, password, status)
		VALUES ($1, $2, $3, $4)`

	_, err := db.QueryContext(
		ctx,
		query,
		a.UserID,
		a.Username,
		a.Password,
		a.Status,
	)
	if err != nil {
		return err
	}
	return nil
}
