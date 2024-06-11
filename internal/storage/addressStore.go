package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreateAddress(address *types.Address) error {
	query := `
	INSERT INTO address (state, city, street, house_number, suburb, postal_code)
	VALUES ($1, $2, $3, $4, $5, $6)`

	res, err := s.db.Query(
		query,
		address.State,
		address.City,
		address.Street,
		address.HouseNumber,
		address.Suburb,
		address.PostalCode,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) UpdateAddress(address *types.Address) error {
	query := `
	UPDATE address
	SET state = ?, city = ?, street = ?, house_number = ?, suburb = ?, postal_code = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		address.State,
		address.City,
		address.Street,
		address.HouseNumber,
		address.Suburb,
		address.PostalCode,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetAddressByEmail(email string) (*types.Address, error) {
	query := `
	SELECT *
	FROM address a
	LEFT JOIN user u
	ON a.id = u.address_id
	WHERE email = $1`

	res, err := s.db.Query(query, email)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		return scanIntoAddress(res)
	}

	return nil, fmt.Errorf("address %s not found", email)
}

func scanIntoAddress(rows *sql.Rows) (*types.Address, error) {
	address := new(types.Address)
	err := rows.Scan(
		&address.ID,
		&address.State,
		&address.City,
		&address.Street,
		&address.HouseNumber,
		&address.Suburb,
		&address.PostalCode,
	)

	if err != nil {
		return nil, err
	}
	return address, err
}
