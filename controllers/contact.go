package controllers

import (
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func (c *Controllers) GetAllContactsInfo(w http.ResponseWriter, r *http.Request) error {
	var contactInfo models.ContactInfo

	var contactInfoList []*models.ContactInfo
	contactInfoList, err := contactInfo.GetAllContactInfo()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, contactInfoList)
	return nil
}
