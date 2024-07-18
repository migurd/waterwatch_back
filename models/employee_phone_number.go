package models

import (
	"context"
	"database/sql"
)

type EmployeePhoneNumber struct {
	EmployeeID  int64  `json:"employee_id"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

func (e *EmployeePhoneNumber) CreateEmployeePhoneNumber(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee_phone_number
		(employee_id, country_code, phone_number)
		VALUES ($1, $2, $3)`

	var err error

	if tx != nil {
		_, err = tx.QueryContext(ctx, query, e.EmployeeID, e.CountryCode, e.PhoneNumber)
	} else {
		_, err = db.QueryContext(ctx, query, e.EmployeeID, e.CountryCode, e.PhoneNumber)
	}
	if err != nil {
		return err
	}

	return nil
}
