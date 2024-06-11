package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreateSAAMaintenance(SAA_maintenance *types.SAAMaintenance) error {
	query := `
	INSERT INTO SAA_maintenance (SAA_id, details, requested_date, done_date)
	VALUES ($1, $2, $3, $4)`

	res, err := s.db.Query(
		query,
		SAA_maintenance.SAAID,
		SAA_maintenance.Details,
		SAA_maintenance.RequestedDate,
		SAA_maintenance.DoneDate,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) UpdateSAAMaintenance(SAA_maintenance *types.SAAMaintenance) error {
	query := `
	UPDATE SAA_maintenance
	SET saa_id = ?, details = ?, requested_date = ?, done_date = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		SAA_maintenance.SAAID,
		SAA_maintenance.Details,
		SAA_maintenance.RequestedDate,
		SAA_maintenance.DoneDate,
		SAA_maintenance.ID,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetSAAMaintenanceByEmail(email string) (*types.SAAMaintenance, error) {
	query := `
	SELECT *
	FROM SAA_maintenance m
	LEFT JOIN "SAA" s
	ON s.SAA_maintenance_id = m.id
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
		return scanIntoSAAMaintenance(res)
	}

	return nil, fmt.Errorf("SAA_maintenance %s not found", email)
}

func (s *PostgresStore) GetSAAMaintenances() ([]*types.SAAMaintenance, error) {
	rows, err := s.db.Query("SELECT * FROM SAA_maintenance")
	if err != nil {
		return nil, err

	}

	SAA_maintenances := []*types.SAAMaintenance{}
	for rows.Next() {
		SAA_maintenance, err := scanIntoSAAMaintenance(rows)

		if err != nil {
			return nil, err
		}

		SAA_maintenances = append(SAA_maintenances, SAA_maintenance)
	}
	return SAA_maintenances, nil
}

func scanIntoSAAMaintenance(rows *sql.Rows) (*types.SAAMaintenance, error) {
	SAA_maintenance := new(types.SAAMaintenance)
	err := rows.Scan(
		&SAA_maintenance.ID,
		&SAA_maintenance.SAAID,
		&SAA_maintenance.Details,
		&SAA_maintenance.RequestedDate,
		&SAA_maintenance.DoneDate,
	)

	if err != nil {
		return nil, err
	}
	return SAA_maintenance, err
}
