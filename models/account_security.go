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

const MaxAttempts = 5

func (as *AccountSecurity) CreateAccountSecurity(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `
    INSERT INTO account_security
    (account_client_id, is_password_encrypted)
    VALUES ($1, $2)`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, as.AccountClientID, as.IsPasswordEncrypted)
	} else {
		_, err = db.ExecContext(ctx, query, as.AccountClientID, as.IsPasswordEncrypted)
	}

	if err != nil {
		return err
	}
	return nil
}

func (as *AccountSecurity) CheckIsPasswordEncrypted(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT check_password_encryption($1);`

	var is_pass_encrypted bool
	err := db.QueryRowContext(ctx, query, username).Scan(&is_pass_encrypted)
	if err != nil {
		return false, err
	}
	return is_pass_encrypted, nil
}
