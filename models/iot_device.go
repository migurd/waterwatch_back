package models

import "context"

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

func (i *IotDevice) UpdateIotDevice() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query :=
		`UPDATE iot_device
		SET status = ?
		WHERE id = ?`

	_, err := db.QueryContext(
		ctx,
		query,
		i.Status,
		i.ID,
	)
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
