package types

type Address struct {
	ID          int    `json:"id"`
	State       string `json:"state"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	Suburb      string `json:"suburb"`
	PostalCode  string `json:"postal_code"`
}