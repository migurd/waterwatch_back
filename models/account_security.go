package models

import (
	"context"
	"time"
)

type AccountSecurity struct {
	AccountUserID           int64     `json:"account_user_id"`
	Attempts                int       `json:"attempts"`
	MaxAttempts             int       `json:"max_attempts"`
	LastAttempt             time.Time `json:"last_attempt"`
	LastTimePasswordChanged time.Time `json:"last_time_password_changed"`
	IsPasswordEncrypted     bool      `json:"is_password_encrypted"`
}

func (a *AccountSecurity) CreateAccountSecurity() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO account_security
		(account_user_id)
		VALUES ($1)`

	_, err := db.QueryContext(
		ctx,
		query,
		a.AccountUserID,
	)
	if err != nil {
		return err
	}
	return nil
}
