package models

import (
	"context"
	"time"
)

type SaaRecord struct {
	ID             int64     `json:"id"`
	SaaID          int64     `json:"saa_id"`
	WaterLevel     float64   `json:"water_level"`
	PhLevel        float64   `json:"ph_level"`
	IsContaminated bool      `json:"is_contaminated"`
	Date           time.Time `json:"date"`
}

func (s *SaaRecord) CreateSaaRecord() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO saa_record
		(saa_id, water_level, ph_level, is_contaminated, date)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := db.QueryContext(
		ctx,
		query,
		s.SaaID,
		s.WaterLevel,
		s.PhLevel,
		s.IsContaminated,
		s.Date,
	)
	if err != nil {
		return err
	}
	return nil
}
