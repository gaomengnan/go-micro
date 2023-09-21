// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"time"
)

type Accounts struct {
	ID        int64     `db:"id" json:"id"`
	Owner     string    `db:"owner" json:"owner"`
	Balance   int64     `db:"balance" json:"balance"`
	Currency  string    `db:"currency" json:"currency"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Entries struct {
	ID        int64 `db:"id" json:"id"`
	AccountID int64 `db:"account_id" json:"account_id"`
	// can be negative
	Amount    int64     `db:"amount" json:"amount"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Transfers struct {
	ID            int64 `db:"id" json:"id"`
	FromAccountID int64 `db:"from_account_id" json:"from_account_id"`
	ToAccountID   int64 `db:"to_account_id" json:"to_account_id"`
	// must be positive
	Amount    int64     `db:"amount" json:"amount"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Users struct {
	Username          string    `db:"username" json:"username"`
	HashedPassword    string    `db:"hashed_password" json:"hashed_password"`
	FullName          string    `db:"full_name" json:"full_name"`
	Email             string    `db:"email" json:"email"`
	PasswordChangedAt time.Time `db:"password_changed_at" json:"password_changed_at"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
}
