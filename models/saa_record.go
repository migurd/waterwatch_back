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
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *SaaRecord) GetAllSaaRecordsBySaaID() ([]*SaaRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT * FROM saa_record WHERE saa_id = $1`

	rows, err := db.QueryContext(ctx, query, s.SaaID)
	if err != nil {
		return nil, err
	}

	var saaRecords []*SaaRecord
	for rows.Next() {
		var saaRecord SaaRecord
		err := rows.Scan(
			&saaRecord.ID,
			&saaRecord.SaaID,
			&saaRecord.WaterLevel,
			&saaRecord.PhLevel,
			&saaRecord.IsContaminated,
			&saaRecord.Date,
		)
		if err != nil {
			return nil, err
		}
		saaRecords = append(saaRecords, &saaRecord)
	}

	return saaRecords, nil
}

func (s *SaaRecord) GetLastSaaRecord() (SaaRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT * FROM saa_record WHERE saa_id = $1 ORDER BY date DESC LIMIT 1`

	var saaRecord SaaRecord
	err := db.QueryRowContext(ctx, query, s.SaaID).Scan(
		&saaRecord.ID,
		&saaRecord.SaaID,
		&saaRecord.WaterLevel,
		&saaRecord.PhLevel,
		&saaRecord.IsContaminated,
		&saaRecord.Date,
	)
	if err != nil {
		return SaaRecord{}, err
	}

	return saaRecord, nil
}
