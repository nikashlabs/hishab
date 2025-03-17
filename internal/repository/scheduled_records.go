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

type ScheduledRecordRepository struct {
	queries *sqlc.Queries
}

func NewScheduledRecordRepository(db *pgxpool.Pool) *ScheduledRecordRepository {
	return &ScheduledRecordRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *ScheduledRecordRepository) CreateScheduledRecord(ctx context.Context, record sqlc.CreateScheduledRecordParams) (sqlc.ScheduledRecord, error) {
	return r.queries.CreateScheduledRecord(ctx, record)
}

func (r *ScheduledRecordRepository) GetScheduledRecordByID(ctx context.Context, id int32) (sqlc.ScheduledRecord, error) {
	return r.queries.GetScheduledRecordByID(ctx, id)
}

func (r *ScheduledRecordRepository) ListScheduledRecordsByAccount(ctx context.Context, accountID pgtype.Int4) ([]sqlc.ScheduledRecord, error) {
	return r.queries.ListScheduledRecordsByAccount(ctx, accountID)
}

func (r *ScheduledRecordRepository) ListScheduledRecordsByStatus(ctx context.Context, status pgtype.Text) ([]sqlc.ScheduledRecord, error) {
	return r.queries.ListScheduledRecordsByStatus(ctx, status)
}

func (r *ScheduledRecordRepository) ListScheduledRecordsByType(ctx context.Context, type_ pgtype.Text) ([]sqlc.ScheduledRecord, error) {
	return r.queries.ListScheduledRecordsByType(ctx, type_)
}

func (r *ScheduledRecordRepository) ListScheduledRecordsByUser(ctx context.Context, userID pgtype.Int4) ([]sqlc.ScheduledRecord, error) {
	return r.queries.ListScheduledRecordsByUser(ctx, userID)
}

func (r *ScheduledRecordRepository) UpdateScheduledRecord(ctx context.Context, record sqlc.UpdateScheduledRecordParams) error {
	return r.queries.UpdateScheduledRecord(ctx, record)
}

func (r *ScheduledRecordRepository) DeleteScheduledRecord(ctx context.Context, id int32) error {
	return r.queries.DeleteScheduledRecord(ctx, id)
}

func (r *ScheduledRecordRepository) DeleteScheduledRecordsByAccount(ctx context.Context, accountID pgtype.Int4) error {
	return r.queries.DeleteScheduledRecordsByAccount(ctx, accountID)
}

func (r *ScheduledRecordRepository) DeleteScheduledRecordsByStatus(ctx context.Context, status pgtype.Text) error {
	return r.queries.DeleteScheduledRecordsByStatus(ctx, status)
}

func (r *ScheduledRecordRepository) DeleteScheduledRecordsByType(ctx context.Context, type_ pgtype.Text) error {
	return r.queries.DeleteScheduledRecordsByType(ctx, type_)
}

func (r *ScheduledRecordRepository) DeleteScheduledRecordsByUser(ctx context.Context, userID pgtype.Int4) error {
	return r.queries.DeleteScheduledRecordsByUser(ctx, userID)
}
