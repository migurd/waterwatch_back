package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func (c *Controllers) GetAllActiveSaaForClient(w http.ResponseWriter, r *http.Request) error {
	var saaDetails models.SaaDetails

	claims, err := GetClaims(r)
	if err != nil {
		return err
	}
	clientID := claims.ID

	var saaDetailsList []*models.SaaDetails
	saaDetailsList, err = saaDetails.GetAllActiveSaaForClient(clientID)
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, saaDetailsList)
	return nil
}

func (*Controllers) GetSaaHeight(w http.ResponseWriter, r *http.Request) error {
	var iot_device models.IotDevice

	err := json.NewDecoder(r.Body).Decode(&iot_device)
	if err != nil {
		return err
	}

	height, err := iot_device.GetHeight()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, height)
	return nil
}
