package controllers

import (
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

var mod models.Models
var client = mod.Client

func GetAllClients(w http.ResponseWriter, r *http.Request) error {
	var clients models.Client
	all, err := clients.GetAllClients()
	if err != nil {
		return err
	}
	helpers.WriteJSON(w, http.StatusOK, all)
	return nil
}
