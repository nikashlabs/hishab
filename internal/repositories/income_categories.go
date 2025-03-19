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

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikashlabs/hishab/internal/database/sqlc"
)

type IncomeCategoryRepository struct {
	queries *sqlc.Queries
}

func NewIncomeCategoryRepository(db *pgxpool.Pool) *IncomeCategoryRepository {
	return &IncomeCategoryRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *IncomeCategoryRepository) CreateIncomeCategory(ctx context.Context, category sqlc.CreateIncomeCategoryParams) (sqlc.IncomeCategory, error) {
	return r.queries.CreateIncomeCategory(ctx, category)
}

func (r *IncomeCategoryRepository) GetIncomeCategoryByID(ctx context.Context, id int32) (sqlc.IncomeCategory, error) {
	return r.queries.GetIncomeCategoryByID(ctx, id)
}

func (r *IncomeCategoryRepository) GetIncomeCategoriesByUserID(ctx context.Context, userID pgtype.Int4) ([]sqlc.IncomeCategory, error) {
	return r.queries.GetIncomeCategoriesByUserID(ctx, userID)
}

func (r *IncomeCategoryRepository) UpdateIncomeCategory(ctx context.Context, category sqlc.UpdateIncomeCategoryParams) error {
	return r.queries.UpdateIncomeCategory(ctx, category)
}

func (r *IncomeCategoryRepository) DeleteIncomeCategory(ctx context.Context, id int32) error {
	return r.queries.DeleteIncomeCategory(ctx, id)
}
