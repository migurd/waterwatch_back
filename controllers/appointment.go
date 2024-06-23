package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/migurd/waterwatch_back/models"
)

// used if the client doesn't have an account, it creates account from scratch and sets up everything needed for a client
func CreateClientAppointment(w http.ResponseWriter, r *http.Request) error {
	var client models.Client
	var client_email models.ClientEmail
	var client_phone_number models.ClientPhoneNumber
	var client_address models.ClientAddress
	var appointment models.Appointment
	var employee models.Employee

	// SAVE UP DATA FROM BODY
	// client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		return err
	}
	err = json.NewDecoder(r.Body).Decode(&client_address)
	if err != nil {
		return err
	}
	err = json.NewDecoder(r.Body).Decode(&client_email)
	if err != nil {
		return err
	}
	err = json.NewDecoder(r.Body).Decode(&client_phone_number)
	if err != nil {
		return err
	}
	err = json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		return err
	}

	// CREATE CLIENT!!!!
	id, err := client.CreateClient()
	if err != nil {
		return err
	}

	client_address.ClientID = id
	client_email.ClientID = id
	client_phone_number.ClientID = id
	err = client_address.CreateClientAddress()
	if err != nil {
		return err
	}
	err = client_email.CreateClientEmail()
	if err != nil {
		return err
	}
	err = client_phone_number.CreateClientPhoneNumber()
	if err != nil {
		return err
	}

	// create appointment
	appointment.ClientID = id
	randomEmployee, err := employee.GetRandomActiveEmployee() // get random employee, might replace later TODO
	if err != nil {
		return err
	}
	appointment.EmployeeID = randomEmployee.ID
	err = appointment.CreateAppointment()
	if err != nil {
		return err
	}

	return nil
}

func CompleteAppointment(w http.ResponseWriter, r *http.Request) error {
	var saa_type models.SaaType
	var saa models.Saa
	var saa_description models.SaaDescription
	var iot_device models.IotDevice
	var client_email models.ClientEmail
	var account models.Account
	var account_security models.AccountSecurity
	var appointment models.Appointment

	err := json.NewDecoder(r.Body).Decode(&saa_type)
	if err != nil {
		return err
	}
	err := json.NewDecoder(r.Body).Decode(&client_email) // id is saa types by now, then replaced
	if err != nil {
		return err
	}

	// get client id
	client_id, err := client_email.GetClientEmailIDByEmail()
	if err != nil {
		return err
	}
	
	// get already existent appointment
	appointment.ClientID = client_id
	err = appointment.GetAppointmentByClientID()
	if err != nil {
		return err
	}

	// create saa type
	saa_type.CreateSaaType()

	// create saa
	saa.

	// create saa description (default name)

	// enable iot device

	// if account doesn't exist, then create acccount and acc security

	return nil
}

// user already has an account
func CreateClientAppointmentWithEmail(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateMaintenanceAppointment(w http.ResponseWriter, r *http.Request) error {
	return nil
}
