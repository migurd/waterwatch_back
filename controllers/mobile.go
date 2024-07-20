package controllers

import (
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func (c *Controllers) GetHome(w http.ResponseWriter, r *http.Request) error {
	var account models.Account
	var homeDetails models.HomeDetails

	claims, err := GetClaims(r)
	if err != nil {
		return err
	}
	account.ClientID = claims.ID

	homeDetails, err = account.GetHomeDetails()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, homeDetails)
	return nil
}
