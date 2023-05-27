package models

import "errors"

var (

	// ErrNoRecord is returned when there is no record returned from the database.
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials is returned when a user tries to login with an
	// incorrect email address or password
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail is returned if a user tries to signup with an email address
	// that is already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
