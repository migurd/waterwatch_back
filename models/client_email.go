package models

import "context"

type ClientEmail struct {
	ClientID int64  `json:"client_id"`
	Email    string `json:"email"`
}

func CreateClientEmail(c *ClientEmail) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO client_email
		(client_id, email)
		VALUES ($1, $2)`

	_, err := db.QueryContext(
		ctx,
		query,
		c.ClientID,
		c.Email,
	)
	if err != nil {
		return err
	}
	return nil
}
