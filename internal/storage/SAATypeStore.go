package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreateSAAType(SAA_type *types.SAAType) error {
	query := `
	INSERT INTO SAA_type (name, description, capacity)
	VALUES ($1, $2, $3)`

	res, err := s.db.Query(
		query,
		SAA_type.Name,
		SAA_type.Description,
		SAA_type.Capacity,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) UpdateSAAType(SAA_type *types.SAAType) error {
	query := `
	UPDATE SAA_type
	SET name = ?, description = ?, capacity = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		SAA_type.Name,
		SAA_type.Description,
		SAA_type.Capacity,
		SAA_type.Id,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetSAATypeByEmail(email string) (*types.SAAType, error) {
	query := `
	SELECT *
	FROM SAA_type  t
	LEFT JOIN "SAA" s
	ON s.SAA_type_id = t.id
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
		return scanIntoSAAType(res)
	}

	return nil, fmt.Errorf("SAA_type %s not found", email)
}

func (s *PostgresStore) GetSAATypes() ([]*types.SAAType, error) {
	rows, err := s.db.Query("SELECT * FROM SAA_type;")
	if err != nil {
		return nil, err

	}

	SAA_types := []*types.SAAType{}
	for rows.Next() {
		SAA_type, err := scanIntoSAAType(rows)

		if err != nil {
			return nil, err
		}

		SAA_types = append(SAA_types, SAA_type)
	}
	return SAA_types, nil
}

func scanIntoSAAType(rows *sql.Rows) (*types.SAAType, error) {
	SAA_type := new(types.SAAType)
	err := rows.Scan(
		&SAA_type.Id,
		&SAA_type.Name,
		&SAA_type.Description,
		&SAA_type.Capacity,
	)

	if err != nil {
		return nil, err
	}
	return SAA_type, err
}


