package models

import (
	"database/sql"
	"time"
)

var db *sql.DB

// timeout of db transaction
const timeoutDB = time.Second * 3

type Models struct {
	// I wonder if this bullcrap will work
	Client              Client
	ClientAddress       ClientAddress
	ClientEmail         ClientEmail
	ClientPhoneNumber   ClientPhoneNumber
	AppointmentType     AppointmentType
	Appointment         Appointment
	Account             Account
	AccountSecurity     AccountSecurity
	Employee            Employee
	EmployeeEmail       EmployeeEmail
	EmployeePhoneNumber EmployeePhoneNumber
	IotDevice           IotDevice
	SaaType             SaaType
	Saa                 Saa
	SaaMaintenance      SaaMaintenance
	SaaRecord           SaaRecord
	SaaDescription      SaaDescription
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}
