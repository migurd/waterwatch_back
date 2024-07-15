package models

import (
	"context"
	"database/sql"
	"time"
)

type AccountSecurity struct {
	AccountClientID         int64     `json:"account_client_id"`
	Attempts                int       `json:"attempts"`
	IsPasswordEncrypted     bool      `json:"is_password_encrypted"`
	LastAttempt             time.Time `json:"last_attempt"`
	LastTimePasswordChanged time.Time `json:"last_time_password_changed"`
}

var MAX_ATTEMPTS = 5

func (as *AccountSecurity) CreateAccountSecurity(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `
    INSERT INTO account_security
    (account_client_id)
    VALUES ($1)`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, as.AccountClientID)
	} else {
		_, err = db.ExecContext(ctx, query, as.AccountClientID)
	}

	if err != nil {
		return err
	}
	return nil
}
