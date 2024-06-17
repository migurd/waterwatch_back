package models

import "context"

type SaaDescription struct {
	SaaID       int64  `json:"saa_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateSaaDescription(s *SaaDescription) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO saa_description
		(saa_id, name, description)
		VALUES ($1, $2, $3)`

	_, err := db.QueryContext(
		ctx,
		query,
		s.SaaID,
		s.Name,
		s.Description,
	)
	if err != nil {
		return err
	}
	return nil
}
