package models

import "context"

type AppointmentType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (c *AppointmentType) CreateAppointmentType() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO appointment_type
		(name)
		VALUES ($1)`

	_, err := db.QueryContext(
		ctx,
		query,
		c.Name,
	)
	if err != nil {
		return err
	}
	return nil
}
