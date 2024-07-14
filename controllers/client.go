package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func CreateClient(w http.ResponseWriter, r *http.Request) error {
	var client models.Client
	
	// get vars from Content-Type
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		return err
	}

	// create client
	client.ID, err = client.CreateClient()
	if err != nil {
		return err
	}

	return nil
}

func CheckClientEmail(w http.ResponseWriter, r *http.Request) error {
	var client_email models.ClientEmail
	err := json.NewDecoder(r.Body).Decode(&client_email)
	if err != nil {
		return err
	}

	is_repeated, err := client_email.CheckClientEmail()
	if err != nil {
		return err
	}
	if is_repeated {
		return fmt.Errorf("correo ya regisrado")
	}

	return nil
}

func GetAllClients(w http.ResponseWriter, r *http.Request) error {
	var clients models.Client
	all, err := clients.GetAllClients()
	if err != nil {
		return err
	}
	helpers.WriteJSON(w, http.StatusOK, all)
	return nil
}
