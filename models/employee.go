package models

import "context"

type Employee struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Curp      string `json:"curp"`
}

func CreateEmployee(e *Employee) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee
		(first_name, last_name, curp)
		VALUES ($1, $2, $3)`

	var id int64
	err := db.QueryRowContext(
		ctx,
		query,
		e.FirstName,
		e.LastName,
		e.Curp,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
