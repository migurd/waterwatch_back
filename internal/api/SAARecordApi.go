package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *APIServer) handleSAARecord(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		if email := r.Header.Get("email"); email != "" {
			return s.handleGetSAARecord(w, r, email)
		}
		return fmt.Errorf("email wasn't found for the SAA_record requested.")
	}
	if r.Method == "POST" {
		return s.handleCreateSAARecord(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetSAARecord(w http.ResponseWriter, r *http.Request, email string) error {
	SAA_record, err := s.store.GetSAARecordByEmail(email)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, SAA_record)
}

func (s *APIServer) handleCreateSAARecord(w http.ResponseWriter, r *http.Request) error {
	createSAARecordReq := new(types.SAARecord)
	if err := json.NewDecoder(r.Body).Decode(createSAARecordReq); err != nil {
		return err
	}

	SAA_record := &types.SAARecord{
		SAAId:          createSAARecordReq.SAAId,
		WaterLevel:     createSAARecordReq.WaterLevel,
		PHLevel:        createSAARecordReq.PHLevel,
		IsContaminated: createSAARecordReq.IsContaminated,
		Date:           createSAARecordReq.Date,
	}
	if err := s.store.CreateSAARecord(SAA_record); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, SAA_record)
}
