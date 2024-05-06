package main

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
	Id         int    `json:"id"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	AddressId  int    `json:"address_id"`
}

type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Account struct {
	Id       *int   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserId   *int   `json:"user_id"`
}

func NewAccount(username string, password string) *Account {
	return &Account{
		Username: username,
		Password: password,
	}
}
