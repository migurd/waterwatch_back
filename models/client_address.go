package models

import (
	"context"
	"database/sql"
)

type ClientAddress struct {
	ID           int64  `json:"id"`
	ClientID     int64  `json:"client_id"`
	State        string `json:"state"`
	City         string `json:"city"`
	Street       string `json:"street"`
	HouseNumber  string `json:"house_number"`
	Neighborhood string `json:"neighborhood"`
	PostalCode   string `json:"postal_code"`
}

func (c *ClientAddress) CreateClientAddress(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `
    INSERT INTO client_address 
    (client_id, state, city, street, house_number, neighborhood, postal_code)
    VALUES ($1, $2, $3, $4, $5, $6, $7)`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, c.ClientID, c.State, c.City, c.Street, c.HouseNumber, c.Neighborhood, c.PostalCode)
	} else {
		_, err = db.ExecContext(ctx, query, c.ClientID, c.State, c.City, c.Street, c.HouseNumber, c.Neighborhood, c.PostalCode)
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *ClientAddress) GetClientAddress() (ClientAddress, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT id, client_id, state, city, street, house_number, neighborhood, postal_code
            FROM client_address WHERE id = $1`

	var clientAddress ClientAddress
	err := db.QueryRowContext(ctx, query, c.ID).Scan(
		&clientAddress.ID,
		&clientAddress.ClientID,
		&clientAddress.State,
		&clientAddress.City,
		&clientAddress.Street,
		&clientAddress.HouseNumber,
		&clientAddress.Neighborhood,
		&clientAddress.PostalCode,
	)
	if err != nil {
		return ClientAddress{}, err
	}
	return clientAddress, nil
}

func (c *ClientAddress) GetAllClientAddress() ([]*ClientAddress, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT id, client_id, state, city, street, house_number, neighborhood, postal_code
            FROM client_address WHERE client_id = $1`

	var clientAddresses []*ClientAddress
	rows, err := db.QueryContext(ctx, query, c.ClientID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var clientAddress ClientAddress
		err = rows.Scan(
			&clientAddress.ID,
			&clientAddress.ClientID,
			&clientAddress.State,
			&clientAddress.City,
			&clientAddress.Street,
			&clientAddress.HouseNumber,
			&clientAddress.Neighborhood,
			&clientAddress.PostalCode,
		)
		if err != nil {
			return nil, err
		}
		clientAddresses = append(clientAddresses, &clientAddress)
	}
	return clientAddresses, nil
}

func (c *ClientAddress) UpdateClientAddress() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `UPDATE client_address SET
            state = $1, city = $2, street = $3, house_number = $4, neighborhood = $5, postal_code = $6
            WHERE id = $7`

	_, err := db.ExecContext(
		ctx,
		query,
		c.State,
		c.City,
		c.Street,
		c.HouseNumber,
		c.Neighborhood,
		c.PostalCode,
		c.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *ClientAddress) DeleteClientAddress() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `DELETE * FROM client_address WHERE id = $1`

	_, err := db.ExecContext(ctx, query, c.ID)
	if err != nil {
		return err
	}
	return nil
}
