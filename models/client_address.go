package models

import "context"

type ClientAddress struct {
	ClientID     int64  `json:"client_id"`
	State        string `json:"state"`
	City         string `json:"city"`
	Street       string `json:"street"`
	HouseNumber  string `json:"house_number"`
	Neighborhood string `json:"neighborhood"`
	PostalCode   string `json:"postal_code"`
}

func (c *ClientAddress) CreateClientAddress() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO client_address 
		(client_id, state, city, street, house_number, neighborhood, postal_code)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.QueryContext(
		ctx,
		query,
		c.ClientID,
		c.State,
		c.City,
		c.Street,
		c.HouseNumber,
		c.Neighborhood,
		c.PostalCode,
	)
	if err != nil {
		return err
	}

	return nil
}
