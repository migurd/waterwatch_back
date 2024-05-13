package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreateSAARecord(SAA_record *types.SAARecord) error {
	query := `
	INSERT INTO SAA_record (name, description, capacity)
	VALUES ($1, $2, $3)`

	res, err := s.db.Query(
		query,
		SAA_record.SAAId,
		SAA_record.WaterLevel,
		SAA_record.PHLevel,
		SAA_record.IsContaminated,
		SAA_record.Date,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) UpdateSAARecord(SAA_record *types.SAARecord) error {
	query := `
	UPDATE SAA_record
	SET saa_id = ?, water_level = ?, pH_level = ?, is_contaminated = ?, date = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		SAA_record.SAAId,
		SAA_record.WaterLevel,
		SAA_record.PHLevel,
		SAA_record.IsContaminated,
		SAA_record.Date,
		SAA_record.Id,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetSAARecordByEmail(email string) (*types.SAARecord, error) {
	query := `
	SELECT *
	FROM SAA_record  t
	LEFT JOIN "SAA" s
	ON s.SAA_record_id = t.id
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
		return scanIntoSAARecord(res)
	}

	return nil, fmt.Errorf("SAA_record %s not found", email)
}

func (s *PostgresStore) GetSAARecords() ([]*types.SAARecord, error) {
	rows, err := s.db.Query("SELECT * FROM SAA_record;")
	if err != nil {
		return nil, err

	}

	SAA_records := []*types.SAARecord{}
	for rows.Next() {
		SAA_record, err := scanIntoSAARecord(rows)

		if err != nil {
			return nil, err
		}

		SAA_records = append(SAA_records, SAA_record)
	}
	return SAA_records, nil
}

func scanIntoSAARecord(rows *sql.Rows) (*types.SAARecord, error) {
	SAA_record := new(types.SAARecord)
	err := rows.Scan(
		&SAA_record.Id,
		&SAA_record.SAAId,
		&SAA_record.WaterLevel,
		&SAA_record.PHLevel,
		&SAA_record.IsContaminated,
		&SAA_record.Date,
	)

	if err != nil {
		return nil, err
	}
	return SAA_record, err
}
