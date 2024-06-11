package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/migurd/waterwatch_back/internal/types"
	"github.com/migurd/waterwatch_back/helpers"
)

func (s *APIServer) handleMicrocontroller(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		if email := r.Header.Get("email"); email != "" {
			return s.handleGetMicrocontroller(w, r, email)
		}
		return fmt.Errorf("email wasn't found for the microcontroller requested")
	}
	if r.Method == "POST" {
		return s.handleCreateMicrocontroller(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetMicrocontroller(w http.ResponseWriter, _ *http.Request, email string) error {
	microcontroller, err := s.store.GetMicrocontrollerByEmail(email)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, microcontroller)
}

func (s *APIServer) handleCreateMicrocontroller(w http.ResponseWriter, r *http.Request) error {
	createMicrocontrollerReq := new(types.Microcontroller)
	if err := json.NewDecoder(r.Body).Decode(createMicrocontrollerReq); err != nil {
		return err
	}

	microcontroller := &types.Microcontroller{
		SerialKey: createMicrocontrollerReq.SerialKey,
		Status:    createMicrocontrollerReq.Status,
	}
	if err := s.store.CreateMicrocontroller(microcontroller); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, microcontroller)
}
