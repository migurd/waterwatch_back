package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/services"
)

type EmployeeAccount struct {
	EmployeeID int64  `json:"employee_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func (e *EmployeeAccount) CreateEmployeeAccount(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO employee_account(employee_id, username, password)
		VALUES (?, ?, ?)`

	var err error

	if tx != nil {
		_, err = tx.QueryContext(ctx, query, e.EmployeeID, e.Username, e.Password)
	} else {
		_, err = db.QueryContext(ctx, query, e.EmployeeID, e.Username, e.Password)
	}
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeAccount) EmployeeLogin() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	// we assume it's possible the user tries to login with email and pass
	query := `SELECT username FROM get_employee_user_by_email($1)`

	var username string
	err := db.QueryRowContext(ctx, query, e.Username).Scan(&username)
	if err == nil {
		e.Username = username // email was sent through username field, so we retrieve the username using email
	} else {
		if err == sql.ErrNoRows {
			// No rows were returned, so we assume an actual username was sent
		} else {
			// An error occurred during query execution
			return "", err
		}
	}

	// we check if the user exists
	query = `SELECT username, password FROM employee_account WHERE username = $1`
	var dbAccount EmployeeAccount

	err = db.QueryRowContext(ctx, query, e.Username).Scan(&dbAccount.Username, &dbAccount.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("username/email not found")
		}
		return "", err
	}

	// Increase attempts
	_, err = db.ExecContext(ctx, `CALL increase_attempts_employee($1)`, e.Username)
	if err != nil {
		return "", err
	}

	// Check if password is encrypted or not
	var employeeAccountSecurity EmployeeAccountSecurity
	is_encrypted, err := employeeAccountSecurity.CheckIsPasswordEncryptedEmployee(e.Username)
	if err != nil {
		return "", err
	}

	// Check if passwords match
	if !is_encrypted {
		if e.Password != dbAccount.Password {
			return "", fmt.Errorf("password is incorrect")
		}
	} else {
		if !services.CheckPasswordHash(e.Password, dbAccount.Password) {
			return "", fmt.Errorf("password is incorrect")
		}
	}

	// It was successfull! reset count of attempts
	_, err = db.ExecContext(ctx, `CALL reset_attempts_employee($1)`, e.Username)
	if err != nil {
		return "", err
	}

	// Generate JWT Token
	token, err := services.GenerateJWT(e.EmployeeID, e.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
