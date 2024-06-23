package models

import "context"

type ClientPhoneNumber struct {
	ClientID    int64  `json:"client_id"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

func (c *ClientPhoneNumber) CreateClientPhoneNumber() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO client_phone_number
		(client_id, country_code, phone_number)
		VALUES ($1, $2, $3)`

	_, err := db.QueryContext(
		ctx,
		query,
		c.ClientID,
		c.CountryCode,
		c.PhoneNumber,
	)
	if err != nil {
		return err
	}
	return nil
}
