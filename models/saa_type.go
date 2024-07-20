package models

import (
	"context"
	"database/sql"
)

type SaaType struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Capacity    int     `json:"capacity"`
	Diameter    float64 `json:"diameter"`
	Height      float64 `json:"height"`
}

func (s *SaaType) CreateSaaType(tx *sql.Tx) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO saa_type
		(name, description, capacity, diameter, height)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	var id int64
	var err error

	if tx != nil {
		err = tx.QueryRowContext(ctx, query, s.Name, s.Description, s.Capacity, s.Diameter, s.Height).Scan(&id)
	} else {
		err = db.QueryRowContext(ctx, query, s.Name, s.Description, s.Capacity, s.Diameter, s.Height).Scan(&id)
	}
	if err != nil {
		return 0, err
	}
	return id, nil
}
