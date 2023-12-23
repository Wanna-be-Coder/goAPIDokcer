// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error)
	CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntries(ctx context.Context, id int64) error
	DeleteTransfers(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetEntries(ctx context.Context, id int64) (Entry, error)
	GetTransfers(ctx context.Context, id int64) (Transfer, error)
	ListAaccount(ctx context.Context, arg ListAaccountParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateEntries(ctx context.Context, arg UpdateEntriesParams) (Entry, error)
	Updatetransfers(ctx context.Context, arg UpdatetransfersParams) (Transfer, error)
}

var _ Querier = (*Queries)(nil)
