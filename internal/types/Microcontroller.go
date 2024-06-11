package types

type Microcontroller struct {
	ID        int    `json:"id"`
	SerialKey string `json:"serial_key"`
	Status    bool   `json:"status"`
}
