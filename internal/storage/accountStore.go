package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreateAccount(account *types.Account) error {
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

func (s *PostgresStore) UpdateAccount(account *types.Account) error {
	query := `
	UPDATE account
	SET username = ?, password = ?, user_id = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		account.Username,
		account.Password,
		account.UserId,
		account.Id,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetAccountByEmail(email string) (*types.Account, error) {
	query := `
	SELECT *
	FROM account a
	LEFT JOIN user u
	ON a.user_id = u.id
	WHERE email = $1`

	res, err := s.db.Query(query, email)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		return scanIntoAccount(res)
	}

	return nil, fmt.Errorf("account %s not found", email)
}

func (s *PostgresStore) GetAccounts() ([]*types.Account, error) {
	rows, err := s.db.Query("SELECT * FROM account;")
	if err != nil {
		return nil, err

	}

	accounts := []*types.Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*types.Account, error) {
	account := new(types.Account)
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
