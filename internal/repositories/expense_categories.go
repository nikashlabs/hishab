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

type ExpenseCategoryRepository struct {
	queries *sqlc.Queries
}

func NewExpenseCategoryRepository(db *pgxpool.Pool) *ExpenseCategoryRepository {
	return &ExpenseCategoryRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *ExpenseCategoryRepository) CreateExpenseCategory(ctx context.Context, category sqlc.CreateExpenseCategoryParams) (sqlc.ExpenseCategory, error) {
	return r.queries.CreateExpenseCategory(ctx, category)
}

func (r *ExpenseCategoryRepository) GetExpenseCategoryByID(ctx context.Context, id int32) (sqlc.ExpenseCategory, error) {
	return r.queries.GetExpenseCategoryByID(ctx, id)
}

func (r *ExpenseCategoryRepository) GetExpenseCategoriesByUserID(ctx context.Context, userID pgtype.Int4) ([]sqlc.ExpenseCategory, error) {
	return r.queries.GetExpenseCategoriesByUserID(ctx, userID)
}

func (r *ExpenseCategoryRepository) UpdateExpenseCategory(ctx context.Context, category sqlc.UpdateExpenseCategoryParams) error {
	return r.queries.UpdateExpenseCategory(ctx, category)
}

func (r *ExpenseCategoryRepository) DeleteExpenseCategory(ctx context.Context, id int32) error {
	return r.queries.DeleteExpenseCategory(ctx, id)
}
