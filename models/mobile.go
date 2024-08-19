package models

import (
	"context"
)

type HomeDetails struct {
	Username        string `json:"username"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	FullPhoneNumber string `json:"full_phone_number"`
}
type SaaDetails struct {
	SaaID                    int64   `json:"saa_id"`
	SerialKey                string  `json:"serial_key"`
	FullAddress              string  `json:"full_address"`
	SaaName                  string  `json:"saa_name"`
	SaaDescription           string  `json:"saa_description"`
	IsGood                   string  `json:"is_good"`
	IsGoodDescription        string  `json:"is_good_description"`
	SaaHeight                int     `json:"saa_height"`
	CurrentSaaCapacity       int     `json:"current_saa_capacity"`
	MaxSaaCapacity           int     `json:"max_saa_capacity"`
	SaaHeight2               int     `json:"saa_height2"`
	CurrentSaaCapacity2      int     `json:"current_saa_capacity2"`
	MaxSaaCapacity2          int     `json:"max_saa_capacity2"`
	DaysSinceLastMaintenance int     `json:"days_since_last_maintenance"`
	WaterLevel               float64 `json:"water_level"`
	WaterLevel2              float64 `json:"water_level2"`
	PhLevel                  float64 `json:"ph_level"`
}

func (a *Account) GetHomeDetails() (HomeDetails, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT username, full_name, email, phone_number FROM get_account_details($1)`

	var homeDetails HomeDetails
	err := db.QueryRowContext(ctx, query, a.ClientID).Scan(
		&homeDetails.Username,
		&homeDetails.FullName,
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
			&saaDetails.SaaName,
			&saaDetails.SaaDescription,
			&saaDetails.IsGood,
			&saaDetails.IsGoodDescription,
			&saaDetails.CurrentSaaCapacity,
			&saaDetails.MaxSaaCapacity,
			&saaDetails.SaaHeight,
			&saaDetails.CurrentSaaCapacity2,
			&saaDetails.MaxSaaCapacity2,
			&saaDetails.SaaHeight2,
			&saaDetails.DaysSinceLastMaintenance,
			&saaDetails.WaterLevel,
			&saaDetails.WaterLevel2,
			&saaDetails.PhLevel,
		)
		if err != nil {
			return nil, err
		}
		saaDetailsList = append(saaDetailsList, &saaDetails)
	}
	return saaDetailsList, nil
}
