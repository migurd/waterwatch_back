package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
	GetAccounts() ([]*Account, error)
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

// func (s *PostgresStore) Init() error {
// }

func (s *PostgresStore) CreateAccount(account *Account) error {
	query := `
	INSERT INTO account (username, password, user_id)
	VALUES ($1, $2, $3)`

	res, err := s.db.Query(query, account.Username, account.Password, account.UserId)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	res, err := s.db.Query("SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		return scanIntoAccount(res)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account;")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.Id,
		&account.Username,
		&account.Password,
		&account.UserId,
	)

	if err != nil {
		return nil, err
	}
	return account, err
}
