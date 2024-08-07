package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func (*Controllers) GetSaaDescription(w http.ResponseWriter, r *http.Request) error {
	var saaDescription models.SaaDescription

	err := json.NewDecoder(r.Body).Decode(&saaDescription)
	if err != nil {
		return err
	}

	saaDescription, err = saaDescription.GetSaaDescription()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, saaDescription)
	return nil
}

func (*Controllers) UpdateSaaDescription(w http.ResponseWriter, r *http.Request) error {
	var saaDescription models.SaaDescription

	err := json.NewDecoder(r.Body).Decode(&saaDescription)
	if err != nil {
		return err
	}

	err = saaDescription.UpdateSaaDescription()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response{Message: "Saa description updated successfully!"})
	return nil
}

