package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/models"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) error {
	var employee models.Employee
	var employee_email models.EmployeeEmail
	var employee_phone_number models.EmployeePhoneNumber
	var employee_account models.EmployeeAccount
	var employee_account_security models.EmployeeAccountSecurity

	// getting data from http
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		return err
	}
	err = json.NewDecoder(r.Body).Decode(&employee_email)
	if err != nil {
		return err
	}
	err = json.NewDecoder(r.Body).Decode(&employee_phone_number)
	if err != nil {
		return err
	}

	// adding to the db
	id, err := employee.CreateEmployee()
	if err != nil {
		return err
	}
	employee_email.EmployeeID = id
	employee_phone_number.EmployeeID = id
	err = employee_email.CreateEmployeeEmail()
	if err != nil {
		return err
	}
	err = employee_phone_number.CreateEmployeePhoneNumber()
	if err != nil {
		return err
	}

	// create account
	employee_account.EmployeeID = id
	employee_account.Username = helpers.GenerateUser(employee.FirstName, employee.Curp) // don't ask me
	// employee_account.Password = helpers.GenerateSecurePassword(10)
	employee_account.Password = "123"
	err = employee_account.CreateEmployeeAccount()
	if err != nil {
		return err
	}
	employee_account_security.EmployeeAccountEmployeeID = id
	err = employee_account_security.CreateEmployeeAccountSecurity()
	if err != nil {
		return err
	}

	return nil
}
