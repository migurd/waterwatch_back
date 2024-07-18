package models

import (
	"context"
	"database/sql"
	"time"
)

type EmployeeAccountSecurity struct {
	EmployeeAccountEmployeeID int64     `json:"employee_account_employee_id"`
	Attempts                  int       `json:"attempts"`
	IsPasswordEncrypted       bool      `json:"is_password_encrypted"`
	LastAttempt               time.Time `json:"last_attempt"`
	LastTimePasswordChanged   time.Time `json:"last_time_password_changed"`
}

func (e *EmployeeAccountSecurity) CreateEmployeeAccountSecurity(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee_account_security
		(employee_account_employee_id)
		VALUES (?)`

	var err error

	if tx != nil {
		_, err = tx.QueryContext(ctx, query, e.EmployeeAccountEmployeeID)
	} else {
		_, err = db.QueryContext(ctx, query, e.EmployeeAccountEmployeeID)
	}
	if err != nil {
		return err
	}

	return nil
}

func (as *EmployeeAccountSecurity) CheckIsPasswordEncryptedEmployee(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT check_password_encryption_employee($1);`

	var is_pass_encrypted bool
	err := db.QueryRowContext(ctx, query, username).Scan(&is_pass_encrypted)
	if err != nil {
		return false, err
	}
	return is_pass_encrypted, nil
}
