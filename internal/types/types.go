package types

import "time"

// Account
type Address struct {
	Id          int    `json:"id"`
	State       string `json:"state"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	Suburb      string `json:"suburb"`
	PostalCode  string `json:"postal_code"`
}

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AddressId int    `json:"address_id"`
}

type Account struct {
	Id       *int   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserId   *int   `json:"user_id"`
}

type AccountSecurity struct {
	Id                      *int      `json:"id"`
	UserId                  int       `json:"user_id"`
	Attempts                int       `json:"attempts"`
	LastAttempt             time.Time `json:"last_attempt"`
	LastTimePasswordChanged time.Time `json:"last_time_password_changed"`
}

type PhoneNumber struct {
	Id          int    `json:"id"`
	AccountId   int    `json:"account_id"`
	PhoneNumber string `json:"phone_number"`
}

// SAA
type Microcontroller struct {
	Id        int    `json:"id"`
	SerialKey string `json:"serial_key"`
	Status    bool   `json:"status"`
}

type SAAType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"` // in lts
}

type SAARecord struct {
	Id             int       `json:"id"`
	SAAId          int       `json:"SAA_id"`
	WaterLevel     int       `json:"water_level"`
	PHLevel        int       `json:"pH_level"`
	IsContaminated bool      `json:"is_contaminated"`
	Date           time.Time `json:"date"`
}

type SAAMaintenance struct {
	Id            int       `json:"id"`
	SAAId         int       `json:"SAA_id"`
	Details       string    `json:"details"`
	RequestedDate time.Time `json:"requested_date"`
	DoneDate      time.Time `json:"done_date"`
}

type SAA struct {
	Id                int `json:"id"`
	SAATypeId         int `json:"SAA_type_id"`
	MicrocontrollerId int `json:"microcontroller_id"`
	AccountId         int `json:"account_id"`
	AddressId         int `json:"address_id"`
}
