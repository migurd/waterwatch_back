package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/services"
)

type Account struct {
	ClientID int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

func (a *Account) CreateAccount(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `
    INSERT INTO account 
    (client_id, username, password, status)
    VALUES ($1, $2, $3, $4)`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, a.ClientID, a.Username, a.Password, a.Status)
	} else {
		_, err = db.ExecContext(ctx, query, a.ClientID, a.Username, a.Password, a.Status)
	}

	if err != nil {
		return err
	}
	return nil
}

func (a *Account) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	// we assume it's possible the user tries to login with email and pass
	query := `SELECT username FROM get_client_user_by_email($1)`

	var username string
	err := db.QueryRowContext(ctx, query, a.Username).Scan(&username)
	if err == nil {
		a.Username = username // email was sent through username field, so we retrieve the username using email
	} else {
		if err == sql.ErrNoRows {
			// No rows were returned, so we assume an actual username was sent
		} else {
			// An error occurred during query execution
			return "", err
		}
	}

	// we check if the user exists
	query = `SELECT username, password FROM account WHERE username = $1`
	var dbAccount Account

	err = db.QueryRowContext(ctx, query, a.Username).Scan(&dbAccount.Username, &dbAccount.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("username/email not found")
		}
		return "", err
	}

	// Increase attempts
	_, err = db.ExecContext(ctx, `CALL increase_attempts_client($1)`, a.Username)
	if err != nil {
		return "", err
	}

	// Check if password is encrypted or not
	var accountSecurity AccountSecurity
	is_encrypted, err := accountSecurity.CheckIsPasswordEncrypted(a.Username)
	if err != nil {
		return "", err
	}

	// Check if passwords match
	if !is_encrypted {
		if a.Password != dbAccount.Password {
			return "", fmt.Errorf("password is incorrect")
		}
	} else {
		if !services.CheckPasswordHash(a.Password, dbAccount.Password) {
			return "", fmt.Errorf("password is incorrect")
		}
	}

	// It was successfull! reset count of attempts
	_, err = db.ExecContext(ctx, `CALL reset_attempts_client($1)`, a.Username)
	if err != nil {
		return "", err
	}

	// Generate JWT Token
	token, err := services.GenerateJWT(a.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Account) DoesAccountExist() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT id FROM account WHERE id = ?`

	var id int64
	row := db.QueryRowContext(ctx, query, a.ClientID)
	err := row.Scan(&id)
	if err != nil { // if it doesn't return any row, then error
		return false, nil
	}
	return true, nil
}
