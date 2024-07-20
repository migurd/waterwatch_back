package models

import (
	"context"
	"database/sql"
)

type SaaDescription struct {
	SaaID       int64  `json:"saa_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *SaaDescription) CreateSaaDescription(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO saa_description
		(saa_id, name, description)
		VALUES ($1, $2, $3)`

	var err error
	if tx != nil {
		_, err = db.ExecContext(ctx, query, s.SaaID, s.Name, s.Description)
	} else {
		_, err = db.ExecContext(ctx, query, s.SaaID, s.Name, s.Description)
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *SaaDescription) GetSaaDescription() (SaaDescription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT * FROM saa_description WHERE saa_id = $1`

	var saaDescription SaaDescription
	err := db.QueryRowContext(ctx, query, s.SaaID).Scan(
		&saaDescription.SaaID,
		&saaDescription.Name,
		&saaDescription.Description,
	)
	if err != nil {
		return SaaDescription{}, err
	}
	return saaDescription, nil
}

func (s *SaaDescription) UpdateSaaDescription() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `UPDATE saa_description SET name = $1, description = $2`

	_, err := db.ExecContext(ctx, query, s.Name, s.Description)
	if err != nil {
		return err
	}

	return nil
}
