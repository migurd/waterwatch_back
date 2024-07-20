package models

import (
	"context"

	"github.com/lib/pq"
)

type Contact struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url"`
}
type ContactEmail struct {
	ContactID int64  `json:"contact_id"`
	Email     string `json:"email"`
}
type ContactPhoneNumber struct {
	ContactID   int64  `json:"contact_id"`
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"PhoneNumber"`
}

type ContactInfo struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	PhotoURL       string   `json:"photo_url"`
	Email          []string `json:"emails"`
	StrPhoneNumber []string `json:"StrPhoneNumbers"`
}

func (c *ContactInfo) GetAllContactInfo() ([]*ContactInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT * FROM get_all_contacts()`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var contactsInfo []*ContactInfo
	for rows.Next() {
		var contactInfo ContactInfo
		err = rows.Scan(
			&contactInfo.ID,
			&contactInfo.Name,
			&contactInfo.PhotoURL,
			pq.Array(&contactInfo.Email),
			pq.Array(&contactInfo.StrPhoneNumber),
		)
		if err != nil {
			return nil, err
		}
		contactsInfo = append(contactsInfo, &contactInfo)
	}
	return contactsInfo, nil
}

func (c *Contact) GetAllContacts() ([]*Contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT id, name, photo_url FROM contact`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var contacts []*Contact
	for rows.Next() {
		var contact Contact
		err = rows.Scan(
			&contact.ID,
			&contact.Name,
			&contact.PhotoURL,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}
	return contacts, nil
}
