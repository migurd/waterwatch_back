package models

import (
	"context"
	"database/sql"
)

type Client struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (c *Client) CreateClient(tx *sql.Tx) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO client
		(first_name, last_name)
		VALUES ($1, $2)
		RETURNING id`

	var id int64
	var err error

	if tx != nil {
		err = tx.QueryRowContext(ctx, query, c.FirstName, c.LastName).Scan(&id)
	} else {
		err = db.QueryRowContext(ctx, query, c.FirstName, c.LastName).Scan(&id)
	}

	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (c *Client) GetAllClients() ([]*Client, error) {
	// Ctx so it closes after 3 scs
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `select * from client`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var clients []*Client
	for rows.Next() { // save each row
		var client Client
		err := rows.Scan( // save each parameter in client
			&client.ID,
			&client.FirstName,
			&client.LastName,
		)
		if err != nil {
			return nil, err
		}
		clients = append(clients, &client) // im speechless by this way of appending
	}
	return clients, nil
}

func (c *Client) GetClientIDByEmail(email string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT * FROM client c
		LEFT JOIN client_email ce
		ON c.id = ce.client_id
		WHERE ce.email = $1`

	var client Client
	row := db.QueryRowContext(
		ctx,
		query,
		email,
	)
	row.Scan(
		&client.ID,
		&client.FirstName,
		&client.LastName,
	)
	return client.ID, nil
}
