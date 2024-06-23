package models

import "context"

type SaaType struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Capacity    int     `json:"capacity"`
	Diameter    float64 `json:"diameter"`
	Height      float64 `json:"height"`
}

func (s *SaaType) CreateSaaType() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO saa_type
		(name, description, capacity, diameter, height)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := db.QueryContext(
		ctx,
		query,
		s.Name,
		s.Description,
		s.Capacity,
		s.Diameter,
		s.Height,
	)
	if err != nil {
		return err
	}
	return nil
}
