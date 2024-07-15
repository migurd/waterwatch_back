package controllers

import (
	"net/http"
)

// used if the client doesn't have an account, it creates account from scratch and sets up everything needed for a client
func CreateClientAppointment(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func CompleteAppointment(w http.ResponseWriter, r *http.Request) error {

	return nil
}

// user already has an account
func CreateClientAppointmentWithEmail(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateMaintenanceAppointment(w http.ResponseWriter, r *http.Request) error {
	return nil
}
