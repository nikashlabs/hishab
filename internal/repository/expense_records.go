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

type ExpenseRecordRepository struct {
	queries *sqlc.Queries
}

func NewExpenseRecordRepository(db *pgxpool.Pool) *ExpenseRecordRepository {
	return &ExpenseRecordRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *ExpenseRecordRepository) CreateExpenseRecord(ctx context.Context, record sqlc.CreateExpenseRecordParams) (sqlc.ExpenseRecord, error) {
	return r.queries.CreateExpenseRecord(ctx, record)
}

func (r *ExpenseRecordRepository) GetExpenseRecordByID(ctx context.Context, id int32) (sqlc.ExpenseRecord, error) {
	return r.queries.GetExpenseRecordByID(ctx, id)
}

func (r *ExpenseRecordRepository) ListExpenseRecordsByAccount(ctx context.Context, accountID pgtype.Int4) ([]sqlc.ExpenseRecord, error) {
	return r.queries.ListExpenseRecordsByAccount(ctx, accountID)
}

func (r *ExpenseRecordRepository) ListExpenseRecordsByCategory(ctx context.Context, categoryID pgtype.Int4) ([]sqlc.ExpenseRecord, error) {
	return r.queries.ListExpenseRecordsByCategory(ctx, categoryID)
}

func (r *ExpenseRecordRepository) ListExpenseRecordsByUser(ctx context.Context, userID pgtype.Int4) ([]sqlc.ExpenseRecord, error) {
	return r.queries.ListExpenseRecordsByUser(ctx, userID)
}

func (r *ExpenseRecordRepository) UpdateExpenseRecord(ctx context.Context, record sqlc.UpdateExpenseRecordParams) error {
	return r.queries.UpdateExpenseRecord(ctx, record)
}

func (r *ExpenseRecordRepository) DeleteExpenseRecord(ctx context.Context, id int32) error {
	return r.queries.DeleteExpenseRecord(ctx, id)
}

func (r *ExpenseRecordRepository) DeleteExpenseRecordsByAccount(ctx context.Context, accountID pgtype.Int4) error {
	return r.queries.DeleteExpenseRecordsByAccount(ctx, accountID)
}

func (r *ExpenseRecordRepository) DeleteExpenseRecordsByCategory(ctx context.Context, categoryID pgtype.Int4) error {
	return r.queries.DeleteExpenseRecordsByCategory(ctx, categoryID)
}

func (r *ExpenseRecordRepository) DeleteExpenseRecordsByUser(ctx context.Context, userID pgtype.Int4) error {
	return r.queries.DeleteExpenseRecordsByUser(ctx, userID)
}
