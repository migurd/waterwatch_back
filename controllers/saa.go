package controllers

import (
	"errors"
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
	// Extract the 'serial_key' from the query parameters
	serialKey := r.URL.Query().Get("serial_key")

	if serialKey == "" {
		return errors.New("serial_key is required")
	}

	// Find the IoT device based on the serial_key
	var iot_device models.IotDevice
	iot_device.SerialKey = serialKey

	// Calculate the height using the IoT device's method
	height, height2, err := iot_device.GetHeight()
	if err != nil {
		return err
	}

	// Prepare the response
	response := map[string]interface{}{
		"height":  height,
		"height2": height2,
	}

	// Send the response as JSON
	helpers.WriteJSON(w, http.StatusOK, response)
	return nil
}
