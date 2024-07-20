package models

import (
	"context"
)

type HomeDetails struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	FullPhoneNumber string `json:"full_phone_number"`
}
type SaaDetails struct {
	SaaID       int64  `json:"saa_id"`
	SerialKey   string `json:"serial_key"`
	FullAddress string `json:"full_address"`
	IsGood      string `json:"is_good"`
}

func (a *Account) GetHomeDetails() (HomeDetails, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT username, email, phone_number FROM get_account_details($1)`

	var homeDetails HomeDetails
	err := db.QueryRowContext(ctx, query, a.ClientID).Scan(
		&homeDetails.Username,
		&homeDetails.Email,
		&homeDetails.FullPhoneNumber,
	)
	if err != nil {
		return HomeDetails{}, err
	}
	return homeDetails, nil
}

func (sd *SaaDetails) GetAllActiveSaaForClient(client_id int64) ([]*SaaDetails, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT * FROM get_all_active_saa_for_client($1)`

	rows, err := db.QueryContext(ctx, query, client_id)
	if err != nil {
		return nil, err
	}

	var saaDetailsList []*SaaDetails
	for rows.Next() {
		var saaDetails SaaDetails
		err = rows.Scan(
			&saaDetails.SaaID,
			&saaDetails.SerialKey,
			&saaDetails.FullAddress,
			&saaDetails.IsGood,
		)
		if err != nil {
			return nil, err
		}
		saaDetailsList = append(saaDetailsList, &saaDetails)
	}
	return saaDetailsList, nil
}
