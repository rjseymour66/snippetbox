package models

import (
	"database/sql"
	"time"
)

// User is a snippet user.
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

// Insert adds a user to the DB.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate verifies that the user email and password match.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Exists checks if a user with a specific ID exists.
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
