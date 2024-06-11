package storage

import (
	"database/sql"
	"fmt"

	"github.com/migurd/waterwatch_back/helpers"
	"github.com/migurd/waterwatch_back/internal/types"
)

func (s *PostgresStore) CreateUser(user *types.User) error {
	var userID int
	query := `
	INSERT INTO "user" (email, first_name, last_name)
	VALUES ($1, $2, $3)
	RETURNING id`

	err := s.db.QueryRow(
		query,
		user.Email,
		user.FirstName,
		user.LastName,
		// user.AddressID,
	).Scan(&userID)

	if err != nil {
		return err
	}

	// Once user has been created, then create account
	pass, err := helpers.GenerateSecurePassword(12)
	if err != nil {
		return err
	}
	username := helpers.GenerateUser(user.FirstName, user.LastName)

	s.CreateAccount(&types.Account{
		UserID:   userID,
		Password: pass,
		Username: username,
	})

	return nil
}

func (s *PostgresStore) UpdateUser(user *types.User) error {
	query := `
	UPDATE user
	SET email = ?, first_name = ?, last_name = ?, address_id = ?
	WHERE id = ?`

	res, err := s.db.Query(
		query,
		user.Email,
		user.FirstName,
		user.LastName,
		user.AddressID,
		user.ID,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) GetUserByEmail(email string) (*types.User, error) {
	res, err := s.db.Query("SELECT * FROM user WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		return scanIntoUser(res)
	}

	return nil, fmt.Errorf("user %s not found", email)
}

func (s *PostgresStore) GetUsers() ([]*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM user;")
	if err != nil {
		return nil, err

	}

	users := []*types.User{}
	for rows.Next() {
		user, err := scanIntoUser(rows)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func scanIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.AddressID,
	)

	if err != nil {
		return nil, err
	}
	return user, err
}
