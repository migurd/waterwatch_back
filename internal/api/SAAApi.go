package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *APIServer) handleSAA(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		if email := r.Header.Get("email"); email != "" {
			return s.handleGetSAA(w, r, email)
		}
		return fmt.Errorf("email wasn't found for the SAA requested")
	}
	if r.Method == "POST" {
		return s.handleCreateSAA(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetSAA(w http.ResponseWriter, _ *http.Request, email string) error {
	SAA_obj, err := s.store.GetSAAByEmail(email)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, SAA_obj)
}

func (s *APIServer) handleCreateSAA(w http.ResponseWriter, r *http.Request) error {
	createSAAReq := new(types.SAA)
	if err := json.NewDecoder(r.Body).Decode(createSAAReq); err != nil {
		return err
	}

	SAA_obj := &types.SAA{
		SAATypeID:         createSAAReq.SAATypeID,
		MicrocontrollerID: createSAAReq.MicrocontrollerID,
		AccountID:         createSAAReq.AccountID,
	}
	if err := s.store.CreateSAA(SAA_obj); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, SAA_obj)
}
