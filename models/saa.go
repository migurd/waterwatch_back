package models

import (
	"context"
	"database/sql"
)

type Saa struct {
	ID            int64 `json:"id"`
	AppointmentID int64 `json:"appointment_id"`
	SaaTypeID     int64 `json:"saa_type_id"`
	IotDeviceID   int64 `json:"iot_device_id"`
}

func (s *Saa) CreateSaa(tx *sql.Tx) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO saa
		(appointment_id, saa_type_id, iot_device_id)
		VALUES ($1, $2, $3)
		RETURNING id`

	var id int64
	var err error

	if tx != nil {
		err = tx.QueryRowContext(ctx, query, s.AppointmentID, s.SaaTypeID, s.IotDeviceID).Scan(&id)
	} else {
		err = db.QueryRowContext(ctx, query, s.AppointmentID, s.SaaTypeID, s.IotDeviceID).Scan(&id)
	}
	if err != nil {
		return 0, err
	}

	return id, nil
}
