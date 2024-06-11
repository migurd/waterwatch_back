package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/migurd/waterwatch_back/internal/types"
	"github.com/migurd/waterwatch_back/helpers"
)

func (s *APIServer) handleSAAType(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		if email := r.Header.Get("email"); email != "" {
			return s.handleGetSAAType(w, r, email)
		}
		return fmt.Errorf("email wasn't found for the SAA_type requested")
	}
	if r.Method == "POST" {
		return s.handleCreateSAAType(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetSAAType(w http.ResponseWriter, _ *http.Request, email string) error {
	SAA_type, err := s.store.GetSAATypeByEmail(email)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, SAA_type)
}

func (s *APIServer) handleCreateSAAType(w http.ResponseWriter, r *http.Request) error {
	createSAATypeReq := new(types.SAAType)
	if err := json.NewDecoder(r.Body).Decode(createSAATypeReq); err != nil {
		return err
	}

	SAA_type := &types.SAAType{
		Name:        createSAATypeReq.Name,
		Description: createSAATypeReq.Description,
		Capacity:    createSAATypeReq.Capacity,
	}
	if err := s.store.CreateSAAType(SAA_type); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, SAA_type)
}
