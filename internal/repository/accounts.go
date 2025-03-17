// Copyright (C) 2025 Nikash Labs
//
// This file is part of hishab.
//
// hishab is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// hishab is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with hishab.  If not, see <https://www.gnu.org/licenses/>.

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikashlabs/hishab/internal/database/sqlc"
)

type AccountRepository struct {
	queries *sqlc.Queries
}

func NewAccountRepository(db *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *AccountRepository) CreateAccount(ctx context.Context, account sqlc.CreateAccountParams) (sqlc.Account, error) {
	return r.queries.CreateAccount(ctx, account)
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, id int32) (sqlc.Account, error) {
	return r.queries.GetAccountByID(ctx, id)
}

func (r *AccountRepository) GetAccountsByUserID(ctx context.Context, userID pgtype.Int4) ([]sqlc.Account, error) {
	return r.queries.GetAccountsByUserID(ctx, userID)
}

func (r *AccountRepository) UpdateAccount(ctx context.Context, account sqlc.UpdateAccountParams) error {
	return r.queries.UpdateAccount(ctx, account)
}

func (r *AccountRepository) DeleteAccount(ctx context.Context, id int32) error {
	return r.queries.DeleteAccount(ctx, id)
}
