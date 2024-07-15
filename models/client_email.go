package models

import (
	"context"
	"database/sql"
)

type ClientEmail struct {
	ClientID int64  `json:"client_id"`
	Email    string `json:"email"`
}

func (ce *ClientEmail) CreateClientEmail(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `
    INSERT INTO client_email
    (client_id, email)
    VALUES ($1, $2)`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, ce.ClientID, ce.Email)
	} else {
		_, err = db.ExecContext(ctx, query, ce.ClientID, ce.Email)
	}

	if err != nil {
		return err
	}
	return nil
}

func (c *ClientEmail) CheckClientEmail() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT check_client_email_exists($1)`

	var is_repeated bool
	err := db.QueryRowContext(ctx, query, c.Email).Scan(&is_repeated)
	if err != nil {
		return false, err
	}
	return is_repeated, nil
}

func (c *ClientEmail) GetClientEmailIDByEmail() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT id FROM client_email WHERE email = ?`

	var id int64
	err := db.QueryRowContext(ctx, query, c.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
