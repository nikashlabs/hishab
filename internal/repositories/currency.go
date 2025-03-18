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

package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikashlabs/hishab/internal/database/sqlc"
)

type CurrencyRepository struct {
	queries *sqlc.Queries
}

func NewCurrencyRepository(db *pgxpool.Pool) *CurrencyRepository {
	return &CurrencyRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *CurrencyRepository) CreateCurrency(ctx context.Context, currency sqlc.CreateCurrencyParams) (sqlc.Currency, error) {
	return r.queries.CreateCurrency(ctx, currency)
}

func (r *CurrencyRepository) GetCurrencyByID(ctx context.Context, id int32) (sqlc.Currency, error) {
	return r.queries.GetCurrencyByID(ctx, id)
}

func (r *CurrencyRepository) GetCurrencyByName(ctx context.Context, name string) (sqlc.Currency, error) {
	return r.queries.GetCurrencyByName(ctx, name)
}

func (r *CurrencyRepository) UpdateCurrency(ctx context.Context, currency sqlc.UpdateCurrencyParams) error {
	return r.queries.UpdateCurrency(ctx, currency)
}

func (r *CurrencyRepository) DeleteCurrency(ctx context.Context, id int32) error {
	return r.queries.DeleteCurrency(ctx, id)
}
