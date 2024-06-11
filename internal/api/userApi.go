package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/migurd/waterwatch_back/internal/types"
	"github.com/migurd/waterwatch_back/helpers"
)

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		if email := r.Header.Get("email"); email != "" {
			return s.handleGetUser(w, r, email)
		}
		return fmt.Errorf("email wasn't found for the user requested")
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, _ *http.Request, email string) error {
	user, err := s.store.GetUserByEmail(email)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	createUserReq := new(types.User)
	if err := json.NewDecoder(r.Body).Decode(createUserReq); err != nil {
		return err
	}

	user := &types.User{
		Email:     createUserReq.Email,
		FirstName: createUserReq.FirstName,
		LastName:  createUserReq.LastName,
		AddressID: createUserReq.AddressID,
	}
	if err := s.store.CreateUser(user); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, user)
}
