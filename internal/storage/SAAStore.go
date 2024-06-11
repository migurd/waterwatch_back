package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreateSAA(SAA *types.SAA) error {
	query := `
	INSERT INTO "SAA" (SAA_type_id, microcontroller_id, account_id, address_id)
	VALUES ($1, $2, $3, $4)`

	res, err := s.db.Query(
		query,
		SAA.SAATypeID,
		SAA.MicrocontrollerID,
		SAA.AccountID,
		SAA.AddressID,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) UpdateSAA(SAA *types.SAA) error {
	query := `
	UPDATE "SAA"
	SET SAA_type_id = ?, microcontroller_id = ?, account_id = ?, address_id = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		SAA.SAATypeID,
		SAA.MicrocontrollerID,
		SAA.AccountID,
		SAA.AddressID,
		SAA.ID,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetSAAByEmail(email string) (*types.SAA, error) {
	query := `
	SELECT *
	FROM "SAA" s
	LEFT JOIN account a
	ON s.account_id
	LEFT JOIN user u
	ON a.user_id = u.id
	WHERE email = $1`

	res, err := s.db.Query(query, email)

	if err != nil {
		return nil, err
	}

	for res.Next() {
		return scanIntoSAA(res)
	}

	return nil, fmt.Errorf("SAA %s not found", email)
}

func (s *PostgresStore) GetSAAs() ([]*types.SAA, error) {
	rows, err := s.db.Query(`SELECT * FROM "SAA"`)
	if err != nil {
		return nil, err

	}

	SAAs := []*types.SAA{}
	for rows.Next() {
		SAA, err := scanIntoSAA(rows)

		if err != nil {
			return nil, err
		}

		SAAs = append(SAAs, SAA)
	}
	return SAAs, nil
}

func scanIntoSAA(rows *sql.Rows) (*types.SAA, error) {
	SAA := new(types.SAA)
	err := rows.Scan(
		&SAA.ID,
		&SAA.SAATypeID,
		&SAA.MicrocontrollerID,
		&SAA.AccountID,
		&SAA.AddressID,
	)

	if err != nil {
		return nil, err
	}
	return SAA, err
}
