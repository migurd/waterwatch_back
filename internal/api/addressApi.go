package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *APIServer) handleAddress(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		if email := r.Header.Get("email"); email != "" {
			return s.handleGetAddress(w, r, email)
		}
		return fmt.Errorf("email wasn't sent or found")
	}
	if r.Method == "POST" {
		return s.handleCreateAddress(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAddress(w http.ResponseWriter, r *http.Request, email string) error {
	address, err := s.store.GetAddressByEmail(email)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, address)
}

func (s *APIServer) handleCreateAddress(w http.ResponseWriter, r *http.Request) error {
	createAddressReq := new(types.Address)
	if err := json.NewDecoder(r.Body).Decode(createAddressReq); err != nil {
		return err
	}

	address := &types.Address{
		State:       createAddressReq.State,
		City:        createAddressReq.City,
		Street:      createAddressReq.Street,
		HouseNumber: createAddressReq.HouseNumber,
		Suburb:      createAddressReq.Suburb,
		PostalCode:  createAddressReq.PostalCode,
	}
	if err := s.store.CreateAddress(address); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, address)
}
