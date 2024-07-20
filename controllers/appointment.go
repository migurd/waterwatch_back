package controllers

import (
	"encoding/json"
	"fmt"
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
		var appoinment models.Appointment

		claims, err := GetClaims(r)
		if err != nil {
			return err
		}

		appoinment.ClientID = claims.ID                // current user
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

func (c *Controllers) UpdateAppointment(appointmentType int64) helpers.ApiFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var appointment models.Appointment

		// should be able to edit request date and details, not other fields
		// ASSUMES data was totally filled
		err := json.NewDecoder(r.Body).Decode(&appointment)
		if err != nil {
			return err
		}

		claims, err := GetClaims(r)
		if err != nil {
			return err
		}

		appointment.ClientID = claims.ID // current user
		appointment.AppointmentTypeID = appointmentType

		err = appointment.UpdateAppointmentByClient()
		if err != nil {
			return err
		}

		helpers.WriteJSON(w, http.StatusOK, helpers.Response{Message: "Updated appointment successfully!"})
		return nil
	}
}

func (c *Controllers) DeleteAppointment(appointmentType int64) helpers.ApiFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var appointment models.Appointment

		claims, err := GetClaims(r)
		if err != nil {
			return err
		}

		appointment.ClientID = claims.ID // current user
		appointment.AppointmentTypeID = appointmentType

		err = appointment.CancelAppointmentClient()
		if err != nil {
			return err
		}

		helpers.WriteJSON(w, http.StatusOK, helpers.Response{Message: "Deleted appointment successfully!"})
		return nil
	}
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
	appointment.EmployeeID = claims.ID

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
		var saaType models.SaaType
		var saa models.Saa
		var saaDescription models.SaaDescription

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
		err = json.Unmarshal(body, &saaType)
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

		// installation process
		if appointmentType == 1 {
			saaTypeID, err := saaType.CreateSaaType(tx)
			if err != nil {
				return err
			}

			saa.SaaTypeID = saaTypeID
			saa.AppointmentID = appoinment.ID
			saa.IotDeviceID, err = iotDevice.GetIotDeviceIDBySerialKey()
			if err != nil {
				return err
			}

			saaID, err := saa.CreateSaa(tx)
			if err != nil {
				return err
			}

			saaDescription.SaaID = saaID
			strNumber := fmt.Sprintf("%d", saaID)
			if err != nil {
				return err
			}
			saaDescription.Name = "Nombre " + strNumber
			saaDescription.Name = "Descripción " + strNumber

			err = saaDescription.CreateSaaDescription(tx)
			if err != nil {
				return err
			}

			iotDevice.Status = true
			if err = iotDevice.UpdateIotDevice(tx); err != nil {
				return err
			}
		} // END IF BLOCK

		// installation and maintenance process
		appoinment.AppointmentTypeID = appointmentType
		if err = appoinment.CompleteAppointment(tx); err != nil {
			return err
		}

		return nil
	}
}

func (c *Controllers) GetAllDoneAppointments(appointmentType int64) helpers.ApiFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var appointment models.Appointment

		claims, err := GetClaims(r)
		if err != nil {
			return err
		}

		appointment.ClientID = claims.ID
		appointment.AppointmentTypeID = appointmentType

		var appointments []*models.Appointment
		appointments, err = appointment.GetAllDoneAppointments()
		if err != nil {
			return err
		}

		helpers.WriteJSON(w, http.StatusOK, appointments)
		return nil
	}
}
