package models

import "context"

type EmployeePhoneNumber struct {
	EmployeeID  int64  `json:"employee_id"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

func CreateEmployeePhoneNumber(e *EmployeePhoneNumber) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee_phone_number
		(employee_id, country_code, phone_number)
		VALUES ($1, $2, $3)`

	_, err := db.QueryContext(
		ctx,
		query,
		e.EmployeeID,
		e.CountryCode,
		e.PhoneNumber,
	)
	if err != nil {
		return err
	}
	return nil
}
