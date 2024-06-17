package models

import "context"

type EmployeeEmail struct {
	EmployeeID int64  `json:"employee_id"`
	Email      string `json:"email"`
}

func CreateEmployeeEmail(e *EmployeeEmail) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee_email
		(employee_id, email)
		VALUES ($1, $2)`

	_, err := db.QueryContext(
		ctx,
		query,
		e.EmployeeID,
		e.Email,
	)
	if err != nil {
		return err
	}
	return nil
}
