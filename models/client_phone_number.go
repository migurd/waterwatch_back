package models

import (
	"context"
	"database/sql"
)

type ClientPhoneNumber struct {
	ClientID    int64  `json:"client_id"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

func (cpn *ClientPhoneNumber) CreateClientPhoneNumber(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `
    INSERT INTO client_phone_number
    (client_id, country_code, phone_number)
    VALUES ($1, $2, $3)`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, cpn.ClientID, cpn.CountryCode, cpn.PhoneNumber)
	} else {
		_, err = db.ExecContext(ctx, query, cpn.ClientID, cpn.CountryCode, cpn.PhoneNumber)
	}

	if err != nil {
		return err
	}
	return nil
}
