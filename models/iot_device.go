package models

import "context"

type IotDevice struct {
	ID        int64  `json:"id"`
	SerialKey string `json:"serial_key"`
	Status    bool   `json:"status"`
}

func CreateIotDevice(i *IotDevice) error {
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
