package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
	"github.com/migurd/waterwatch_back/services"
)

func (c *Controllers) CreateClient(w http.ResponseWriter, r *http.Request) error {
	var client models.Client
	var clientAddress models.ClientAddress
	var clientEmail models.ClientEmail
	var clientPhoneNumber models.ClientPhoneNumber
	var account models.Account
	var accountSecurity models.AccountSecurity

	// read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// get vars from Content-Type
	err = json.Unmarshal(body, &client)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &clientAddress)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &clientEmail)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &clientPhoneNumber)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &account)
	if err != nil {
		return err
	}

	// validation
	if len(client.FirstName) == 0 || len(client.LastName) == 0 ||
		len(clientEmail.Email) == 0 || len(clientPhoneNumber.CountryCode) == 0 || len(clientPhoneNumber.PhoneNumber) == 0 ||
		len(clientAddress.City) == 0 || len(clientAddress.HouseNumber) == 0 || len(clientAddress.State) == 0 ||
		len(clientAddress.State) == 0 || len(clientAddress.Neighborhood) == 0 || len(clientAddress.PostalCode) == 0 {
		return fmt.Errorf("rellena todos los campos para crear una cuenta, por favor")
	}

	// hash the password
	hashedPassword, err := services.HashPassword(account.Password)
	if err != nil {
		return err
	}
	account.Password = hashedPassword
	accountSecurity.IsPasswordEncrypted = true

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Defer a function to handle rollback and commit
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // Re-throw panic after rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Create client within the transaction
	client.ID, err = client.CreateClient(tx)
	if err != nil {
		return err
	}

	clientAddress.ClientID = client.ID
	clientEmail.ClientID = client.ID
	clientPhoneNumber.ClientID = client.ID
	account.ClientID = client.ID
	accountSecurity.AccountClientID = client.ID

	if err = clientAddress.CreateClientAddress(tx); err != nil {
		return err
	}
	if err = clientEmail.CreateClientEmail(tx); err != nil {
		return err
	}
	if err = clientPhoneNumber.CreateClientPhoneNumber(tx); err != nil {
		return err
	}
	if err = account.CreateAccount(tx); err != nil {
		return err
	}
	if err = accountSecurity.CreateAccountSecurity(tx); err != nil {
		return err
	}

	// Generate JWT after successful transaction commit
	token, err := services.GenerateJWT(account.Username)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Client %s created!", account.Username)
	helpers.WriteJSON(w, http.StatusCreated, helpers.Response{Message: message, Token: token})
	return nil
}

func (c *Controllers) ClientLogin(w http.ResponseWriter, r *http.Request) error {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		return err
	}

	// if no error, then password and user correct
	token, err := account.Login()
	if err != nil {
		return err
	}

	// Optionally, set the token as an HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	message := fmt.Sprintf("Started session as %s successfully.", account.Username)
	helpers.WriteJSON(w, http.StatusOK, helpers.Response{Message: message, Token: token})
	return nil
}

func (c *Controllers) CheckClientEmail(w http.ResponseWriter, r *http.Request) error {
	var clientEmail models.ClientEmail
	err := json.NewDecoder(r.Body).Decode(&clientEmail)
	if err != nil {
		return err
	}

	if len(clientEmail.Email) == 0 {
		return fmt.Errorf("email empty")
	}

	is_repeated, err := clientEmail.CheckClientEmail()
	if err != nil {
		return err
	}
	if is_repeated {
		helpers.WriteJSON(w, http.StatusOK, true)
	}

	helpers.WriteJSON(w, http.StatusOK, false)
	return nil
}

func (c *Controllers) GetAllClients(w http.ResponseWriter, r *http.Request) error {
	var clients models.Client
	all, err := clients.GetAllClients()
	if err != nil {
		return err
	}
	helpers.WriteJSON(w, http.StatusOK, all)
	return nil
}
