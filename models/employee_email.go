package models

import (
	"context"
	"database/sql"
	"errors"
)

type EmployeeEmail struct {
	EmployeeID int64  `json:"employee_id"`
	Email      string `json:"email"`
}

func (e *EmployeeEmail) CreateEmployeeEmail(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee_email
		(employee_id, email)
		VALUES ($1, $2)`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, e.EmployeeID, e.Email)
	} else {
		_, err = db.ExecContext(ctx, query, e.EmployeeID, e.Email)
	}
	if err != nil {
		return errors.New("error creating employee email: " + err.Error())
	}

	return nil
}
