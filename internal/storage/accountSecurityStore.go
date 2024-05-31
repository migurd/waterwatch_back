package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreateAccountSecurity(account *types.AccountSecurity) error {
	query := `
	INSERT INTO account (user_id, attempts, last_attempt, last_time_password_changed)
	VALUES ($1, $2, $3, $4)`

	res, err := s.db.Query(query, account.UserId, account.Attempts, account.LastAttempt, account.LastTimePasswordChanged)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) UpdateAccountSecurity(account *types.AccountSecurity) error {
	query := `
	UPDATE account
	SET user_id = ?, attempts = ?, last_attempt = ?, last_time_password_changed = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		account.UserId,
		account.Attempts,
		account.LastAttempt,
		account.LastTimePasswordChanged,
		account.Id,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetAccountSecurityByEmail(email string) (*types.AccountSecurity, error) {
	query := `
	SELECT *
	FROM account_security s
	LEFT JOIN account a
	ON s.account_id = a.id
	LEFT JOIN user u
	ON a.user_id = u.id
	WHERE email = $1`

	res, err := s.db.Query(query, email)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		return scanIntoAccountSecurity(res)
	}

	return nil, fmt.Errorf("account %s not found", email)
}

func scanIntoAccountSecurity(rows *sql.Rows) (*types.AccountSecurity, error) {
	account := new(types.AccountSecurity)
	err := rows.Scan(
		&account.Id,
		&account.UserId,
		&account.Attempts,
		&account.LastAttempt,
		&account.LastTimePasswordChanged,
	)

	if err != nil {
		return nil, err
	}
	return account, err
}
