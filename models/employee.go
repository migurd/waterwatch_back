package models

import (
	"context"
	"database/sql"

	"github.com/migurd/waterwatch_back/helpers"
)

type Employee struct {
	ID             int64  `json:"id"`
	EmployeeTypeID int64  `json:"employee_type_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Curp           string `json:"curp"`
	Status         bool   `json:"status"`
}

func scanEmployee(rows *sql.Rows, e *Employee) error {
	rows.Scan(
		&e.ID,
		&e.EmployeeTypeID,
		&e.FirstName,
		&e.LastName,
		&e.Curp,
		&e.Status,
	)
	return nil
}

func (e *Employee) CreateEmployee() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee
		(employee_type_id, first_name, last_name, curp)
		VALUES ($1, $2, $3, $4)`

	var id int64
	err := db.QueryRowContext(
		ctx,
		query,
		e.EmployeeTypeID,
		e.FirstName,
		e.LastName,
		e.Curp,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (e *Employee) GetAllEmployees() ([]*Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT * FROM employee;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, nil
	}

	var employees []*Employee
	for rows.Next() {
		var employee Employee
		err := scanEmployee(rows, &employee)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &employee)
	}
	return employees, nil
}

func (e *Employee) GetAllActiveEmployees() ([]*Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT * FROM employee WHERE status = TRUE`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var employees []*Employee
	for rows.Next() {
		var employee Employee
		err := scanEmployee(rows, &employee)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &employee)
	}
	return employees, nil
}

func (e *Employee) GetEmployeeIDByEmail(email string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT * FROM employee e
		LEFT JOIN employee_email ee
		ON e.id = ee.employee_id
		WHERE ee.email = $1`

	var employee Employee
	row := db.QueryRowContext(
		ctx,
		query,
		email,
	)
	err := row.Scan(
		&employee.ID,
		&employee.EmployeeTypeID,
		&employee.FirstName,
		&employee.LastName,
		&employee.Curp,
		&employee.Status,
	)
	if err != nil {
		return 0, err
	}
	return employee.ID, nil
}

func (e *Employee) GetRandomActiveEmployee() (*Employee, error) {
	var employees []*Employee
	employees, err := e.GetAllActiveEmployees()
	if err != nil {
		return nil, err
	}

	employee := employees[helpers.RandomNumber(len(employees))]
	return employee, nil
}