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

type InstallmentRepository struct {
	queries *sqlc.Queries
}

func NewInstallmentRepository(db *pgxpool.Pool) *InstallmentRepository {
	return &InstallmentRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *InstallmentRepository) CreateInstallment(ctx context.Context, installment sqlc.CreateInstallmentParams) (sqlc.Installment, error) {
	return r.queries.CreateInstallment(ctx, installment)
}

func (r *InstallmentRepository) GetInstallmentByID(ctx context.Context, id int32) (sqlc.Installment, error) {
	return r.queries.GetInstallmentByID(ctx, id)
}

func (r *InstallmentRepository) ListInstallmentsByAccount(ctx context.Context, accountID pgtype.Int4) ([]sqlc.Installment, error) {
	return r.queries.ListInstallmentsByAccount(ctx, accountID)
}

func (r *InstallmentRepository) ListInstallmentsByLoan(ctx context.Context, loanID pgtype.Int4) ([]sqlc.Installment, error) {
	return r.queries.ListInstallmentsByLoan(ctx, loanID)
}

func (r *InstallmentRepository) ListInstallmentsByScheduledRecord(ctx context.Context, scheduledRecordID pgtype.Int4) ([]sqlc.Installment, error) {
	return r.queries.ListInstallmentsByScheduledRecord(ctx, scheduledRecordID)
}

func (r *InstallmentRepository) ListInstallmentsByUser(ctx context.Context, userID pgtype.Int4) ([]sqlc.Installment, error) {
	return r.queries.ListInstallmentsByUser(ctx, userID)
}

func (r *InstallmentRepository) UpdateInstallment(ctx context.Context, installment sqlc.UpdateInstallmentParams) error {
	return r.queries.UpdateInstallment(ctx, installment)
}

func (r *InstallmentRepository) DeleteInstallment(ctx context.Context, id int32) error {
	return r.queries.DeleteInstallment(ctx, id)
}

func (r *InstallmentRepository) DeleteInstallmentsByAccount(ctx context.Context, accountID pgtype.Int4) error {
	return r.queries.DeleteInstallmentsByAccount(ctx, accountID)
}

func (r *InstallmentRepository) DeleteInstallmentsByLoan(ctx context.Context, loanID pgtype.Int4) error {
	return r.queries.DeleteInstallmentsByLoan(ctx, loanID)
}

func (r *InstallmentRepository) DeleteInstallmentsByScheduledRecord(ctx context.Context, scheduledRecordID pgtype.Int4) error {
	return r.queries.DeleteInstallmentsByScheduledRecord(ctx, scheduledRecordID)
}

func (r *InstallmentRepository) DeleteInstallmentsByUser(ctx context.Context, userID pgtype.Int4) error {
	return r.queries.DeleteInstallmentsByUser(ctx, userID)
}
