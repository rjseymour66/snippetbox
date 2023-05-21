package models

import "errors"

// ErrNoRecord is returned when there is no record returned from the database.
var ErrNoRecord = errors.New("models: no matching record found")
