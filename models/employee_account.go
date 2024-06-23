package models

import "context"

type EmployeeAccount struct {
	EmployeeID int64  `json:"employee_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func (e *EmployeeAccount) CreateEmployeeAccount() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee_account(employee_id, username, password)
		VALUES (?, ?, ?)`

	_, err := db.QueryContext(
		ctx,
		query,
		e.EmployeeID,
		e.Username,
		e.Password,
	)
	if err != nil {
		return err
	}

	return nil
}
