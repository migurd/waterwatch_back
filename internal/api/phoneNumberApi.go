package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/migurd/waterwatch_back/internal/types"
	"github.com/migurd/waterwatch_back/helpers"
)

func (s *APIServer) handlePhoneNumber(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		if email := r.Header.Get("email"); email != "" {
			return s.handleGetPhoneNumber(w, r, email)
		}
		return fmt.Errorf("email wasn't found for the phone_number requested")
	}
	if r.Method == "POST" {
		return s.handleCreatePhoneNumber(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetPhoneNumber(w http.ResponseWriter, _ *http.Request, email string) error {
	phone_number, err := s.store.GetPhoneNumberByEmail(email)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, phone_number)
}

func (s *APIServer) handleCreatePhoneNumber(w http.ResponseWriter, r *http.Request) error {
	createPhoneNumberReq := new(types.PhoneNumber)
	if err := json.NewDecoder(r.Body).Decode(createPhoneNumberReq); err != nil {
		return err
	}

	phone_number := &types.PhoneNumber{
		AccountID:   createPhoneNumberReq.AccountID,
		PhoneNumber: createPhoneNumberReq.PhoneNumber,
	}
	if err := s.store.CreatePhoneNumber(phone_number); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, phone_number)
}
