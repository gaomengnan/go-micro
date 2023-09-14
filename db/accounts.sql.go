// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: accounts.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING id, owner, balance, currency, created_at
`

type CreateAccountParams struct {
	Owner    string `db:"owner" json:"owner"`
	Balance  int64  `db:"balance" json:"balance"`
	Currency string `db:"currency" json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
