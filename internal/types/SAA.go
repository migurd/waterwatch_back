package types
import "time"

type SAAType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Capacity    int    `json:"capacity"` // in lts
}

type SAARecord struct {
	ID             int       `json:"id"`
	SAAID          int       `json:"SAA_id"`
	WaterLevel     int       `json:"water_level"`
	PHLevel        int       `json:"pH_level"`
	IsContaminated bool      `json:"is_contaminated"`
	Date           time.Time `json:"date"`
}

type SAAMaintenance struct {
	ID            int       `json:"id"`
	SAAID         int       `json:"SAA_id"`
	Details       string    `json:"details"`
	RequestedDate time.Time `json:"requested_date"`
	DoneDate      time.Time `json:"done_date"`
}

type SAA struct {
	ID                int `json:"id"`
	SAATypeID         int `json:"SAA_type_id"`
	MicrocontrollerID int `json:"microcontroller_id"`
	AccountID         int `json:"account_id"`
	AddressID         int `json:"address_id"`
}