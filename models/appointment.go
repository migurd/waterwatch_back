package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Appointment struct {
	ID                int64     `json:"id"`
	AddressID         int64     `json:"address_id"`
	AppointmentTypeID int64     `json:"appointment_type_id"`
	ClientID          int64     `json:"client_id"`
	EmployeeID        int64     `json:"employee_id"`
	Details           string    `json:"details"`
	RequestedDate     time.Time `json:"requested_date"`
	DoneDate          time.Time `json:"done_date"`
}

func (a *Appointment) CreateAppointment() (int64, error) {
	is_appointment, err := a.IsAppointment()
	if err != nil {
		return 0, err
	}

	if is_appointment {
		return 0, fmt.Errorf("there's an already on-going appointment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `INSERT INTO appointment
		(appointment_type_id, address_id, client_id, details, requested_date)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id`

	var id int64
	err = db.QueryRowContext(
		ctx,
		query,
		a.AppointmentTypeID,
		a.AddressID,
		a.ClientID,
		a.Details,
		a.RequestedDate,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *Appointment) GetPendingAppointment() (Appointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	// TODO: instead of returning the id, it should return the data like an address string or the employee details
	query := `SELECT id, appointment_type_id, address_id, client_id, COALESCE(employee_id, 0), details, requested_date FROM appointment WHERE appointment_type_id = $1 AND done_date IS NULL AND client_id = $2`

	var appointment Appointment
	err := db.QueryRowContext(ctx, query, a.AppointmentTypeID, a.ClientID).Scan(
		&appointment.ID,
		&appointment.AppointmentTypeID,
		&appointment.AddressID,
		&appointment.ClientID,
		&appointment.EmployeeID,
		&appointment.Details,
		&appointment.RequestedDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Appointment{}, fmt.Errorf("there is not an appointment yet")
		}
		return Appointment{}, err
	}

	return appointment, nil
}

func (a *Appointment) UpdateAppointmentByClient() error {
	currentAppointment, err := a.GetPendingAppointment()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `UPDATE appointment SET details = $1, requested_date = $2 WHERE id = $3`

	_, err = db.ExecContext(ctx, query, a.Details, a.RequestedDate, currentAppointment.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *Appointment) UpdateDoneDateAppoinment() error {
	currentAppointment, err := a.GetPendingAppointment()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `UPDATE appointment SET done_date = $1 WHERE id = $2`

	_, err = db.ExecContext(ctx, query, a.DoneDate, currentAppointment.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *Appointment) CancelAppointmentClient() error {
	currentAppointment, err := a.GetPendingAppointment()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `DELETE from appointment WHERE id = $1`

	_, err = db.ExecContext(ctx, query, currentAppointment.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *Appointment) GetAllUnassignedAppointments() ([]*Appointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT id, address_id, appointment_type_id, client_id, employee_id, details, requested_date, done_date
		FROM appointment
		WHERE employee_id IS NULL AND done_date IS NULL`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var appointments []*Appointment
	for rows.Next() {
		var appoinment Appointment
		err = rows.Scan(
			&appoinment.ID,
			&appoinment.AddressID,
			&appoinment.AppointmentTypeID,
			&appoinment.ClientID,
			&appoinment.EmployeeID,
			&appoinment.Details,
			&appoinment.RequestedDate,
			&appoinment.DoneDate,
		)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, &appoinment)
	}

	return appointments, nil
}

func (a *Appointment) GetAllAppoinmentsAssigned() ([]*Appointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT id, address_id, appointment_type_id, client_id, employee_id, details, requested_date, done_date
		FROM appointment
		WHERE employee_id = $1 AND done_date IS NULL`

	rows, err := db.QueryContext(ctx, query, a.EmployeeID)
	if err != nil {
		return nil, err
	}

	var appointments []*Appointment
	for rows.Next() {
		var appoinment Appointment
		err = rows.Scan(
			&appoinment.ID,
			&appoinment.AddressID,
			&appoinment.AppointmentTypeID,
			&appoinment.ClientID,
			&appoinment.EmployeeID,
			&appoinment.Details,
			&appoinment.RequestedDate,
			&appoinment.DoneDate,
		)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, &appoinment)
	}

	return appointments, nil
}

func (a *Appointment) AcceptAppointment() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `UPDATE appointment SET employee_id = $1 WHERE id = $2`

	_, err := db.ExecContext(ctx, query, a.EmployeeID, a.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *Appointment) CancelAppointmentEmployee() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `UPDATE appointment SET employee_id = NULL WHERE id = $1`

	_, err := db.ExecContext(ctx, query, a.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *Appointment) CompleteAppointment(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `UPDATE appointment SET done_date = $1 WHERE id = $2 AND appointment_type_id = $3`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, time.Now(), a.ID, a.AppointmentTypeID)
	} else {
		_, err = db.ExecContext(ctx, query, time.Now(), a.ID, a.AppointmentTypeID)
	}
	if err != nil {
		return err
	}

	return nil
}

func (a *Appointment) IsAppointment() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT id FROM appointment WHERE appointment_type_id = $1 AND done_date IS NULL AND client_id = $2`

	var id int64
	err := db.QueryRowContext(ctx, query, a.AppointmentTypeID, a.ClientID).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
