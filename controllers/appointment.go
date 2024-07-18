package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func (c *Controllers) CreateAppointment(appointmentType int64) helpers.ApiFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var appointment models.Appointment

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, &appointment)
		if err != nil {
			return err
		}

		claims, err := GetClaims(r)
		if err != nil {
			return err
		}

		appointment.ClientID = claims.ID
		appointment.AppointmentTypeID = appointmentType

		_, err = appointment.CreateAppointment()
		if err != nil {
			return err
		}

		helpers.WriteJSON(w, http.StatusCreated, helpers.Response{Message: "Appointment created successfully!"})
		return nil
	}
}

func (c *Controllers) GetPendingAppointment(appointmentType int64) helpers.ApiFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var clientEmail models.ClientEmail
		var appoinment models.Appointment

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, &clientEmail)
		if err != nil {
			return err
		}

		client_id, err := clientEmail.GetClientIDByEmail()
		if err != nil {
			return err
		}
		appoinment.ClientID = client_id                // current user
		appoinment.AppointmentTypeID = appointmentType // installation

		var pendingAppointment models.Appointment
		pendingAppointment, err = appoinment.GetPendingAppointment()
		if err != nil {
			return err
		}

		helpers.WriteJSON(w, http.StatusOK, pendingAppointment)
		return nil
	}
}

func (c *Controllers) UpdateAppointment(w http.ResponseWriter, r *http.Request) error {
	var appointment models.Appointment

	// should be able to edit request date and details, not other fields
	// ASSUMES data was totally filled
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		return err
	}

	err = appointment.UpdateAppointmentByClient()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response{Message: "Updated appointment successfully!"})
	return nil
}

func (c *Controllers) DeleteAppointment(w http.ResponseWriter, r *http.Request) error {
	var appointment models.Appointment

	// ASSUMES ID and appointment_type_id were filled
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		return err
	}

	err = appointment.CancelAppointmentClient()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response{Message: "Deleted appointment successfully!"})
	return nil
}

func (c *Controllers) GetAllUnassignedAppointments(w http.ResponseWriter, r *http.Request) error {
	var appointment models.Appointment

	var appointments []*models.Appointment
	appointments, err := appointment.GetAllUnassignedAppointments()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, appointments)
	return nil
}

func (c *Controllers) GetAllAppointmentsAssigned(w http.ResponseWriter, r *http.Request) error {
	var appointment models.Appointment

	claims, err := GetClaims(r)
	if err != nil {
		return err
	}
	appointment.EmployeeID = claims.ID

	var appointments []*models.Appointment
	appointments, err = appointment.GetAllAppoinmentsAssigned()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, appointments)
	return nil
}

func (c *Controllers) AcceptAppointment(w http.ResponseWriter, r *http.Request) error {
	var appointment models.Appointment

	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		return err
	}

	claims, err := GetClaims(r)
	if err != nil {
		return err
	}
	appointment.ClientID = claims.ID

	err = appointment.AcceptAppointment()
	if err != nil {
		return err
	}

	return nil
}

func (c *Controllers) CancelAppointmentEmployee(w http.ResponseWriter, r *http.Request) error {
	var appointment models.Appointment

	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		return err
	}

	claims, err := GetClaims(r)
	if err != nil {
		return err
	}
	appointment.ClientID = claims.ID

	err = appointment.CancelAppointmentEmployee()
	if err != nil {
		return err
	}

	return nil
}

func (c *Controllers) CompleteAppointment(appointmentType int64) helpers.ApiFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var appoinment models.Appointment
		var iotDevice models.IotDevice

		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, &appoinment)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, &iotDevice)
		if err != nil {
			return err
		}

		// Start a transaction
		tx, err := db.Begin()
		if err != nil {
			return err
		}

		// Defer a function to handle rollback and commit
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r) // Re-throw panic after rollback
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()

		iotDevice.Status = true
		if err = iotDevice.UpdateIotDevice(tx); err != nil {
			return err
		}

		appoinment.AppointmentTypeID = appointmentType
		if err = appoinment.CompleteAppointment(tx); err != nil {
			return err
		}

		return nil
	}
}
