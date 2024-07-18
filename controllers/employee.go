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

func (c *Controllers) CreateEmployee(w http.ResponseWriter, r *http.Request) error {
	var employee models.Employee
	var employeeEmail models.EmployeeEmail
	var employeePhoneNumber models.EmployeePhoneNumber
	var employeeAccount models.EmployeeAccount
	var employeeAccountSecurity models.EmployeeAccountSecurity

	// read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// get vars from Content-Type
	err = json.Unmarshal(body, &employee)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &employeeEmail)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &employeePhoneNumber)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &employeeAccount)
	if err != nil {
		return err
	}

	// validation
	if len(employee.FirstName) == 0 || len(employee.LastName) == 0 ||
		len(employeeEmail.Email) == 0 || len(employeePhoneNumber.CountryCode) == 0 || len(employeePhoneNumber.PhoneNumber) == 0 {
		return fmt.Errorf("rellena todos los campos para crear una cuenta, por favor")
	}

	// hash the password
	hashedPassword, err := services.HashPassword(employeeAccount.Password)
	if err != nil {
		return err
	}
	employeeAccount.Password = hashedPassword
	employeeAccountSecurity.IsPasswordEncrypted = true

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

	// Create employee within the transaction
	employee.ID, err = employee.CreateEmployee(tx)
	if err != nil {
		return err
	}

	employeeEmail.EmployeeID = employee.ID
	employeePhoneNumber.EmployeeID = employee.ID
	employeeAccount.EmployeeID = employee.ID
	employeeAccountSecurity.EmployeeAccountEmployeeID = employee.ID

	if err = employeeEmail.CreateEmployeeEmail(tx); err != nil {
		return err
	}
	if err = employeePhoneNumber.CreateEmployeePhoneNumber(tx); err != nil {
		return err
	}
	if err = employeeAccount.CreateEmployeeAccount(tx); err != nil {
		return err
	}
	if err = employeeAccountSecurity.CreateEmployeeAccountSecurity(tx); err != nil {
		return err
	}

	// Generate JWT after successful transaction commit
	token, err := services.GenerateJWT(employee.ID, employeeAccount.Username)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Employee %s created!", employeeAccount.Username)
	helpers.WriteJSON(w, http.StatusCreated, helpers.Response{Message: message, Token: token})
	return nil
}

func (c *Controllers) EmployeeLogin(w http.ResponseWriter, r *http.Request) error {
	var employeeAccount models.EmployeeAccount
	err := json.NewDecoder(r.Body).Decode(&employeeAccount)
	if err != nil {
		return err
	}

	// if no error, then password and user correct
	token, err := employeeAccount.EmployeeLogin()
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

	message := fmt.Sprintf("Started session as %s successfully.", employeeAccount.Username)
	helpers.WriteJSON(w, http.StatusOK, helpers.Response{Message: message, Token: token})
	return nil
}
