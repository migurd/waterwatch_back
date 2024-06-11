package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreatePhoneNumber(phoneNumber *types.PhoneNumber) error {
	query := `
	INSERT INTO phone_number (account_id, phone_number)
	VALUES ($1, $2)`

	res, err := s.db.Query(
		query,
		phoneNumber.AccountID,
		phoneNumber.PhoneNumber,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) UpdatePhoneNumber(phoneNumber *types.PhoneNumber) error {
	query := `
	UPDATE phone_number
	SET account_id = ?, phone_number = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		phoneNumber.AccountID,
		phoneNumber.PhoneNumber,
		phoneNumber.ID,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetPhoneNumberByEmail(email string) (*types.PhoneNumber, error) {
	query := `
	SELECT *
	FROM phone_number p
	LEFT JOIN account a
	ON p.account_id = a.id
	LEFT JOIN user u
	ON a.user_id = u.id
	WHERE email = $1`

	res, err := s.db.Query(query, email)

	if err != nil {
		return nil, err
	}

	for res.Next() {
		return scanIntoPhoneNumber(res)
	}

	return nil, fmt.Errorf("phone number %s not found", email)
}

func (s *PostgresStore) GetPhoneNumbers() ([]*types.PhoneNumber, error) {
	rows, err := s.db.Query("SELECT * FROM phone_number")
	if err != nil {
		return nil, err

	}

	phoneNumbers := []*types.PhoneNumber{}
	for rows.Next() {
		phoneNumber, err := scanIntoPhoneNumber(rows)

		if err != nil {
			return nil, err
		}

		phoneNumbers = append(phoneNumbers, phoneNumber)
	}
	return phoneNumbers, nil
}

func scanIntoPhoneNumber(rows *sql.Rows) (*types.PhoneNumber, error) {
	phoneNumber := new(types.PhoneNumber)
	err := rows.Scan(
		&phoneNumber.ID,
		&phoneNumber.AccountID,
		&phoneNumber.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}
	return phoneNumber, err
}
