package models

import (
	"context"
	"time"
)

type Appointment struct {
	ID                int64     `json:"id"`
	AppointmentTypeID int64     `json:"appointment_type_id"`
	ClientID          int64     `json:"client_id"`
	EmployeeID        int64     `json:"employee_id"`
	Details           string    `json:"details"`
	RequestedDate     time.Time `json:"requested_date"`
	DoneDate          time.Time `json:"done_date"`
}

func (a *Appointment) CreateAppointment() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO appointment 
		(id, appointment_type_id, client_id, employee_id, details, requested_date, done_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.QueryContext(
		ctx,
		query,
		a.ID,
		a.AppointmentTypeID,
		a.ClientID,
		a.EmployeeID,
		a.Details,
		a.RequestedDate,
		a.DoneDate,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *Appointment) UpdateAppointment() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`UPDATE appointment
		SET employee_id = ?, details = ?, requested_date = ?, done_date = ?
		WHERE id = ?`

	_, err := db.QueryContext(
		ctx,
		query,
		a.EmployeeID,
		a.Details,
		a.RequestedDate,
		a.DoneDate,
		a.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *Appointment) GetAppointmentByClientID() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT * FROM appointment WHERE client_id = ? AND done_date IS NULL`

	row := db.QueryRowContext(ctx, query, a.ClientID)
	row.Scan(
		&a.ID,
		&a.AppointmentTypeID,
		&a.ClientID,
		&a.EmployeeID,
		&a.Details,
		&a.RequestedDate,
		&a.DoneDate,
	)
	return nil
}
