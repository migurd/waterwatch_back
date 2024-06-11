package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/migurd/waterwatch_back/internal/types"
	"github.com/migurd/waterwatch_back/helpers"
)

func (s *APIServer) handleSAAMaintenance(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		if email := r.Header.Get("email"); email != "" {
			return s.handleGetSAAMaintenance(w, r, email)
		}
		return fmt.Errorf("email wasn't found for the SAA_maintenance requested")
	}
	if r.Method == "POST" {
		return s.handleCreateSAAMaintenance(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetSAAMaintenance(w http.ResponseWriter, _ *http.Request, email string) error {
	SAA_maintenance, err := s.store.GetSAAMaintenanceByEmail(email)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, SAA_maintenance)
}

func (s *APIServer) handleCreateSAAMaintenance(w http.ResponseWriter, r *http.Request) error {
	createSAAMaintenanceReq := new(types.SAAMaintenance)
	if err := json.NewDecoder(r.Body).Decode(createSAAMaintenanceReq); err != nil {
		return err
	}

	SAA_maintenance := &types.SAAMaintenance{
		SAAID:         createSAAMaintenanceReq.SAAID,
		Details:       createSAAMaintenanceReq.Details,
		RequestedDate: createSAAMaintenanceReq.RequestedDate,
		DoneDate:      createSAAMaintenanceReq.DoneDate,
	}
	if err := s.store.CreateSAAMaintenance(SAA_maintenance); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, SAA_maintenance)
}
