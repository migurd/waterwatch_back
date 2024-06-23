package models

import (
	"context"
	"time"
)

type EmployeeAccountSecurity struct {
	EmployeeAccountEmployeeID int64     `json:"employee_account_employee_id"`
	Attempts                  int       `json:"attempts"`
	IsPasswordEncrypted       bool      `json:"is_password_encrypted"`
	LastAttempt               time.Time `json:"last_attempt"`
	LastTimePasswordChanged   time.Time `json:"last_time_password_changed"`
}

func (e *EmployeeAccountSecurity) CreateEmployeeAccountSecurity() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee_account_security
		(employee_account_employee_id)
		VALUES (?)`

	_, err := db.QueryContext(ctx, query, e.EmployeeAccountEmployeeID)
	if err != nil {
		return err
	}
	return nil
}