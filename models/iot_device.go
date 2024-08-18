package models

import (
	"context"
	"database/sql"
)

type IotDevice struct {
	ID        int64  `json:"id"`
	SerialKey string `json:"serial_key"`
	Status    bool   `json:"status"`
}

func (i *IotDevice) CreateIotDevice() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`INSERT INTO iot_device
		(serial_key)
		VALUES ($1)`

	_, err := db.QueryContext(
		ctx,
		query,
		i.SerialKey,
	)
	if err != nil {
		return err
	}
	return nil
}

func (i *IotDevice) UpdateIotDevice(tx *sql.Tx) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`UPDATE iot_device
		SET status = $1
		WHERE serial_key = $2`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, i.Status, i.SerialKey)
	} else {
		_, err = db.ExecContext(ctx, query, i.Status, i.SerialKey)
	}
	if err != nil {
		return err
	}

	return nil
}

func (i *IotDevice) GetIotDeviceIDBySerialKey() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT id FROM iot_device WHERE serial_key = ?`

	var id int64
	err := db.QueryRowContext(ctx, query, i.SerialKey).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (i *IotDevice) GetIotDeviceStatus() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT status FROM iot_device WHERE serial_key = $1`

	var status bool
	err := db.QueryRowContext(ctx, query, i.SerialKey).Scan(&status)
	if err != nil {
		return false, err
	}

	return status, nil
}

func (i *IotDevice) GetSaaIDBySerialKey() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `SELECT get_saa_id_by_serial_key($1)`

	var saa_id int64
	err := db.QueryRowContext(ctx, query, i.SerialKey).Scan(&saa_id)
	if err != nil {
		return 0, err
	}

	return saa_id, nil
}

func (i *IotDevice) IsBusy() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`SELECT * FROM saa WHERE serial_key_id = ?`

	var id int64
	row := db.QueryRowContext(ctx, query, i.ID)
	err := row.Scan(&id)
	if err != nil {
		return false, err // row not found
	}
	return true, nil // row found
}

func (i *IotDevice) GetHeight() (int64, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `
	SELECT st.height, st2.height
	FROM iot_device i
	LEFT JOIN saa sa
		ON i.id = sa.iot_device_id
	LEFT JOIN saa_type st
		ON sa.saa_type_id = st.id
	LEFT JOIN saa_type st2
		ON sa.saa_type_id2 = st2.id
	WHERE
		i.serial_key = $1
	`

	var height, height2 int64
	row := db.QueryRowContext(ctx, query, i.SerialKey)
	err := row.Scan(&height, &height2)
	if err != nil {
		return 0, 0, err
	}
	return height, height2, nil
}
