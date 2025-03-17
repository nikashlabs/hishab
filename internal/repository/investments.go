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

type InvestmentRepository struct {
	queries *sqlc.Queries
}

func NewInvestmentRepository(db *pgxpool.Pool) *InvestmentRepository {
	return &InvestmentRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *InvestmentRepository) CreateInvestment(ctx context.Context, investment sqlc.CreateInvestmentParams) (sqlc.Investment, error) {
	return r.queries.CreateInvestment(ctx, investment)
}

func (r *InvestmentRepository) GetInvestmentByID(ctx context.Context, id int32) (sqlc.Investment, error) {
	return r.queries.GetInvestmentByID(ctx, id)
}

func (r *InvestmentRepository) ListInvestmentsByAccount(ctx context.Context, accountID pgtype.Int4) ([]sqlc.Investment, error) {
	return r.queries.ListInvestmentsByAccount(ctx, accountID)
}

func (r *InvestmentRepository) ListInvestmentsByScheduledRecord(ctx context.Context, scheduledRecordID pgtype.Int4) ([]sqlc.Investment, error) {
	return r.queries.ListInvestmentsByScheduledRecord(ctx, scheduledRecordID)
}

func (r *InvestmentRepository) ListInvestmentsByType(ctx context.Context, investmentTypeID pgtype.Int4) ([]sqlc.Investment, error) {
	return r.queries.ListInvestmentsByType(ctx, investmentTypeID)
}

func (r *InvestmentRepository) ListInvestmentsByUser(ctx context.Context, userID pgtype.Int4) ([]sqlc.Investment, error) {
	return r.queries.ListInvestmentsByUser(ctx, userID)
}

func (r *InvestmentRepository) UpdateInvestment(ctx context.Context, investment sqlc.UpdateInvestmentParams) error {
	return r.queries.UpdateInvestment(ctx, investment)
}

func (r *InvestmentRepository) DeleteInvestment(ctx context.Context, id int32) error {
	return r.queries.DeleteInvestment(ctx, id)
}

func (r *InvestmentRepository) DeleteInvestmentsByAccount(ctx context.Context, accountID pgtype.Int4) error {
	return r.queries.DeleteInvestmentsByAccount(ctx, accountID)
}

func (r *InvestmentRepository) DeleteInvestmentsByScheduledRecord(ctx context.Context, scheduledRecordID pgtype.Int4) error {
	return r.queries.DeleteInvestmentsByScheduledRecord(ctx, scheduledRecordID)
}

func (r *InvestmentRepository) DeleteInvestmentsByType(ctx context.Context, investmentTypeID pgtype.Int4) error {
	return r.queries.DeleteInvestmentsByType(ctx, investmentTypeID)
}

func (r *InvestmentRepository) DeleteInvestmentsByUser(ctx context.Context, userID pgtype.Int4) error {
	return r.queries.DeleteInvestmentsByUser(ctx, userID)
}
