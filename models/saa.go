package models

import "context"

type Saa struct {
	ID          int64 `json:"id"`
	ClientID    int64 `json:"client_id"`
	SaaTypeID   int64 `json:"saa_type_id"`
	IotDeviceID int64 `json:"iot_device_id"`
}

func CreateSaa(s *Saa) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO saa
		(client_id, saa_type_id, iot_device_id)
		VALUES ($1, $2, $3)`

	_, err := db.QueryContext(
		ctx,
		query,
		s.ClientID,
		s.SaaTypeID,
		s.IotDeviceID,
	)
	if err != nil {
		return err
	}
	return nil
}
