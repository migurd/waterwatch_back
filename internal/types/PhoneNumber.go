package types

type PhoneNumber struct {
	ID          int    `json:"id"`
	AccountID   int    `json:"account_id"`
	PhoneNumber string `json:"phone_number"`
}