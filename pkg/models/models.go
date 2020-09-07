package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord displayed when no matching records found.
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials error will indicate if a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail error will indicate if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet contains fields from snippets table.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User struct aligning with our User table.
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
