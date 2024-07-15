package controllers

import (
	"encoding/json"
	"errors"
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
	id, err := client.CreateClient(nil)
	if err != nil {
		return err
	}

	client_address.ClientID = id
	client_email.ClientID = id
	client_phone_number.ClientID = id
	err = client_address.CreateClientAddress(nil)
	if err != nil {
		return err
	}
	err = client_email.CreateClientEmail(nil)
	if err != nil {
		return err
	}
	err = client_phone_number.CreateClientPhoneNumber(nil)
	if err != nil {
		return err
	}

	// choose employee to attend the client in an appoinment
	appointment.ClientID = id
	randomEmployee, err := employee.GetRandomActiveEmployee() // get random employee, might replace later TODO
	if err != nil {
		return err
	}
	appointment.EmployeeID = randomEmployee.ID

	// send email to employee chosen with the appointment info

	// create appointment
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
	err = json.NewDecoder(r.Body).Decode(&client_email) // client_email.id holds saa_type id. BEWARE
	if err != nil {
		return err
	}
	err = json.NewDecoder(r.Body).Decode(&iot_device) // iot_device.id holds saa_type id. BEWARE
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

	// enable iot device
	// if serial key is found, then proceed
	iot_device_id, err := iot_device.GetIotDeviceIDBySerialKey()
	if err != nil {
		return err
	}
	iot_device.ID = iot_device_id
	is_iot_device_busy, err := iot_device.IsBusy()
	if err != nil {
		return err
	}
	if is_iot_device_busy {
		return errors.New("IoT device busy")
	}
	iot_device.Status = true
	iot_device.UpdateIotDevice()

	// create saa type
	saa_type.CreateSaaType()

	// create saa
	saa.ClientID = client_id
	saa.IotDeviceID = iot_device_id
	saa.SaaTypeID = saa_type.ID

	// create saa description (default name)
	saa_description.Name = "Nombre por defecto"
	saa_description.Description = "Descripción por defecto"
	saa_description.CreateSaaDescription()

	// if account doesn't exist, then create acccount and acc security
	account.ClientID = client_id
	does_account_exist, err := account.DoesAccountExist()
	if err != nil {
		return err
	}
	if !does_account_exist {
		// create account and account security
		account.CreateAccount(nil)
		account_security.AccountClientID = client_id
		account_security.CreateAccountSecurity(nil)

		// send email to client of account created
		// pray
	}
	return nil
}

// user already has an account
func CreateClientAppointmentWithEmail(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateMaintenanceAppointment(w http.ResponseWriter, r *http.Request) error {
	return nil
}
