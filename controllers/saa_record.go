package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func (c *Controllers) CreateSaaRecord(w http.ResponseWriter, r *http.Request) error {
	var iotDevice models.IotDevice
	var saaRecord models.SaaRecord

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &iotDevice)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &saaRecord)
	if err != nil {
		return err
	}

	iotDeviceStatus, err := iotDevice.GetIotDeviceStatus()
	if err != nil {
		return err
	}
	if !iotDeviceStatus {
		return fmt.Errorf("iot device not enabled. can't post new info")
	}

	saaID, err := iotDevice.GetSaaIDBySerialKey()
	if err != nil {
		return err
	}
	saaRecord.SaaID = saaID

	err = saaRecord.CreateSaaRecord()
	if err != nil {
		return nil
	}

	helpers.WriteJSON(w, http.StatusCreated, helpers.Response{Message: "Created new record!"})
	return nil
}

func (c *Controllers) GetSaaRecords(w http.ResponseWriter, r *http.Request) error {
	var saa models.Saa
	var saaRecord models.SaaRecord // used to extract all records

	err := json.NewDecoder(r.Body).Decode(&saa) // gets id
	if err != nil {
		return err
	}
	saaRecord.SaaID = saa.ID

	saaRecords, err := saaRecord.GetAllSaaRecordsBySaaID()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, saaRecords)
	return nil
}

func (c *Controllers) GetLastSaaRecord(w http.ResponseWriter, r *http.Request) error {
	var saaRecord models.SaaRecord // used to extract all records

	saaID := r.URL.Query().Get("saa_id")
	if saaID == "" {
		// Handle the empty string case
		log.Println("Parameter is empty")
	}
	saaIDValue, err := strconv.ParseInt(saaID, 10, 64)
	if err != nil {
		return err
	}
	saaRecord.SaaID = saaIDValue

	saaRecord, err = saaRecord.GetLastSaaRecord()
	if err != nil {
		return err
	}

	helpers.WriteJSON(w, http.StatusOK, saaRecord)
	return nil

}
