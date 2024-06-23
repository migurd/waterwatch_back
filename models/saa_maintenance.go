package models

import "context"

type SaaMaintenance struct {
	ID            int64 `json:"id"`
	AppointmentID int64 `json:"appointment_id"`
	SaaID         int64 `json:"saa_id"`
}

func (s *SaaMaintenance) CreateSaaMaintenance() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO saa_maintenance
		(appointment_id, saa_id)
		VALUES ($1, $2)`

	_, err := db.QueryContext(
		ctx,
		query,
		s.AppointmentID,
		s.SaaID,
	)
	if err != nil {
		return err
	}
	return nil
}
