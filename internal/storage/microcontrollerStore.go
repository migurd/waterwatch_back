package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreateMicrocontroller(microcontroller *types.Microcontroller) error {
	query := `
	INSERT INTO microcontroller (serial_key, status)
	VALUES ($1, $2)`

	res, err := s.db.Query(
		query,
		microcontroller.SerialKey,
		microcontroller.Status,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) UpdateMicrocontroller(microcontroller *types.Microcontroller) error {
	query := `
	UPDATE microcontroller
	SET serial_key = ?, status = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		microcontroller.SerialKey,
		microcontroller.Status,
		microcontroller.ID,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetMicrocontrollerByEmail(email string) (*types.Microcontroller, error) {
	query := `
	SELECT *
	FROM microcontroller m
	LEFT JOIN SAA s
	ON s.microcontroller_id = m.id
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
		return scanIntoMicrocontroller(res)
	}

	return nil, fmt.Errorf("microcontroller %s not found", email)
}

func (s *PostgresStore) GetMicrocontrollers() ([]*types.Microcontroller, error) {
	rows, err := s.db.Query("SELECT * FROM microcontroller;")
	if err != nil {
		return nil, err

	}

	microcontrollers := []*types.Microcontroller{}
	for rows.Next() {
		microcontroller, err := scanIntoMicrocontroller(rows)

		if err != nil {
			return nil, err
		}

		microcontrollers = append(microcontrollers, microcontroller)
	}
	return microcontrollers, nil
}

func scanIntoMicrocontroller(rows *sql.Rows) (*types.Microcontroller, error) {
	microcontroller := new(types.Microcontroller)
	err := rows.Scan(
		&microcontroller.ID,
		&microcontroller.SerialKey,
		&microcontroller.Status,
	)

	if err != nil {
		return nil, err
	}
	return microcontroller, err
}
