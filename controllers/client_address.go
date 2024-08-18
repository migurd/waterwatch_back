package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func (c *Controllers) CreateClientAddress(w http.ResponseWriter, r *http.Request) error {
	var clientAddress models.ClientAddress

	err := json.NewDecoder(r.Body).Decode(&clientAddress)
	if err != nil {
		return err
	}

	claims, err := GetClaims(r)
	if err != nil {
		return err
	}
	clientAddress.ClientID = claims.ID

	err = clientAddress.CreateClientAddress(nil)
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusCreated, helpers.Response{Message: "Address created successfully!"})
	return nil
}

func (c *Controllers) GetClientAddress(w http.ResponseWriter, r *http.Request) error {
	var clientAddress models.ClientAddress

	claims, err := GetClaims(r)
	if err != nil {
		return err
	}
	clientAddress.ClientID = claims.ID

	clientAddress, err = clientAddress.GetClientAddress()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, clientAddress)
	return nil
}

func (c *Controllers) GetAllClientAddresses(w http.ResponseWriter, r *http.Request) error {
	var clientAddress models.ClientAddress

	claims, err := GetClaims(r)
	if err != nil {
		return err
	}
	clientAddress.ClientID = claims.ID

	var clientAddresses []*models.ClientAddress
	clientAddresses, err = clientAddress.GetAllClientAddress()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, clientAddresses)
	return nil
}

func (c *Controllers) UpdateClientAddress(w http.ResponseWriter, r *http.Request) error {
	var clientAddress models.ClientAddress

	err := json.NewDecoder(r.Body).Decode(&clientAddress)
	if err != nil {
		return err
	}

	claims, err := GetClaims(r)
	if err != nil {
		return err
	}
	clientAddress.ClientID = claims.ID

	err = clientAddress.UpdateClientAddress()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Response{Message: "Address updated successfully!"})
	return nil
}

func (c *Controllers) DeleteClientAddress(w http.ResponseWriter, r *http.Request) error {
	// Extract address_id from the request headers
	addressID := r.Header.Get("address_id")
	if addressID == "" {
		http.Error(w, "Missing address_id in headers", http.StatusBadRequest)
		return fmt.Errorf("missing address_id in headers")
	}
	addressIDValue, err := strconv.ParseInt(addressID, 10, 64)
	if err != nil {
		return err
	}

	// Extract claims from the request
	claims, err := GetClaims(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return err
	}

	// Create a ClientAddress object and set the relevant fields
	clientAddress := models.ClientAddress{
		ID:       addressIDValue,
		ClientID: claims.ID,
	}

	// Attempt to delete the client address
	err = clientAddress.DeleteClientAddress()
	if err != nil {
		http.Error(w, "Failed to delete address", http.StatusInternalServerError)
		return err
	}

	helpers.WriteJSON(w, http.StatusCreated, helpers.Response{Message: "Address deleted successfully!"})
	return nil
}
