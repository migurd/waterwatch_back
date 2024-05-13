package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/migurd/waterwatch_back/internal/types"
)

type Storage interface {
	CreateAddress(*types.Address) error
	UpdateAddress(*types.Address) error
	GetAddressByEmail(string) (*types.Address, error)

	CreateUser(*types.User) error
	UpdateUser(*types.User) error
	GetUserByEmail(string) (*types.User, error)

	CreateAccount(*types.Account) error
	UpdateAccount(*types.Account) error
	GetAccountByEmail(string) (*types.Account, error)
	GetAccounts() ([]*types.Account, error)

	CreatePhoneNumber(*types.PhoneNumber) error
	UpdatePhoneNumber(*types.PhoneNumber) error
	GetPhoneNumberByEmail(string) (*types.PhoneNumber, error)

	CreateMicrocontroller(*types.Microcontroller) error
	UpdateMicrocontroller(*types.Microcontroller) error
	GetMicrocontrollerByEmail(string) (*types.Microcontroller, error)

	CreateSAAType(*types.SAAType) error
	UpdateSAAType(*types.SAAType) error
	GetSAATypeByEmail(string) (*types.SAAType, error)

	CreateSAAMaintenance(*types.SAAMaintenance) error
	UpdateSAAMaintenance(*types.SAAMaintenance) error
	GetSAAMaintenanceByEmail(string) (*types.SAAMaintenance, error)

	CreateSAARecord(*types.SAARecord) error
	UpdateSAARecord(*types.SAARecord) error
	GetSAARecordByEmail(string) (*types.SAARecord, error)

	CreateSAA(*types.SAA) error
	UpdateSAA(*types.SAA) error
	GetSAAByEmail(string) (*types.SAA, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=angelq password=cisco123 dbname=waterwatch sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}
