package models

import "time"

// Queue represents the 'queue' table in the database.
type Queue struct {
	ID          int       `json:"id" db:"id"`
	Transaction string    `json:"transaction" db:"transaction"`
	Wallet      string    `json:"wallet" db:"wallet"`
	Amount      int       `json:"amount" db:"amount"`
	Message     string    `json:"message" db:"message"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Success represents the 'success' table in the database.
type Success struct {
	ID          int       `json:"id" db:"id"`
	Transaction string    `json:"transaction" db:"transaction"`
	Wallet      string    `json:"wallet" db:"wallet"`
	Amount      int       `json:"amount" db:"amount"`
	Message     string    `json:"message" db:"message"`
	Hash        string    `json:"hash" db:"hash"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
