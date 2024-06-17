package controllers

import (
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func CreateClient(w http.ResponseWriter, r *http.Request) error {
	var client models.Client
	_, err := client.CreateClient()
	if err != nil {
		return err
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
