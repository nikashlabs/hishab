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

type IncomeRecordRepository struct {
	queries *sqlc.Queries
}

func NewIncomeRecordRepository(db *pgxpool.Pool) *IncomeRecordRepository {
	return &IncomeRecordRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *IncomeRecordRepository) CreateIncomeRecord(ctx context.Context, record sqlc.CreateIncomeRecordParams) (sqlc.IncomeRecord, error) {
	return r.queries.CreateIncomeRecord(ctx, record)
}

func (r *IncomeRecordRepository) GetIncomeRecordByID(ctx context.Context, id int32) (sqlc.IncomeRecord, error) {
	return r.queries.GetIncomeRecordByID(ctx, id)
}

func (r *IncomeRecordRepository) ListIncomeRecordsByAccount(ctx context.Context, accountID pgtype.Int4) ([]sqlc.IncomeRecord, error) {
	return r.queries.ListIncomeRecordsByAccount(ctx, accountID)
}

func (r *IncomeRecordRepository) ListIncomeRecordsByCategory(ctx context.Context, categoryID pgtype.Int4) ([]sqlc.IncomeRecord, error) {
	return r.queries.ListIncomeRecordsByCategory(ctx, categoryID)
}

func (r *IncomeRecordRepository) ListIncomeRecordsByUser(ctx context.Context, userID pgtype.Int4) ([]sqlc.IncomeRecord, error) {
	return r.queries.ListIncomeRecordsByUser(ctx, userID)
}

func (r *IncomeRecordRepository) UpdateIncomeRecord(ctx context.Context, record sqlc.UpdateIncomeRecordParams) error {
	return r.queries.UpdateIncomeRecord(ctx, record)
}

func (r *IncomeRecordRepository) DeleteIncomeRecord(ctx context.Context, id int32) error {
	return r.queries.DeleteIncomeRecord(ctx, id)
}

func (r *IncomeRecordRepository) DeleteIncomeRecordsByAccount(ctx context.Context, accountID pgtype.Int4) error {
	return r.queries.DeleteIncomeRecordsByAccount(ctx, accountID)
}

func (r *IncomeRecordRepository) DeleteIncomeRecordsByCategory(ctx context.Context, categoryID pgtype.Int4) error {
	return r.queries.DeleteIncomeRecordsByCategory(ctx, categoryID)
}

func (r *IncomeRecordRepository) DeleteIncomeRecordsByUser(ctx context.Context, userID pgtype.Int4) error {
	return r.queries.DeleteIncomeRecordsByUser(ctx, userID)
}
