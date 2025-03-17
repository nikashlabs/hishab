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

type LoanRepository struct {
	queries *sqlc.Queries
}

func NewLoanRepository(db *pgxpool.Pool) *LoanRepository {
	return &LoanRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *LoanRepository) CreateLoan(ctx context.Context, loan sqlc.CreateLoanParams) (sqlc.Loan, error) {
	return r.queries.CreateLoan(ctx, loan)
}

func (r *LoanRepository) GetLoanByID(ctx context.Context, id int32) (sqlc.Loan, error) {
	return r.queries.GetLoanByID(ctx, id)
}

func (r *LoanRepository) ListLoansByAccount(ctx context.Context, accountID pgtype.Int4) ([]sqlc.Loan, error) {
	return r.queries.ListLoansByAccount(ctx, accountID)
}

func (r *LoanRepository) ListLoansByType(ctx context.Context, loanTypeID pgtype.Int4) ([]sqlc.Loan, error) {
	return r.queries.ListLoansByType(ctx, loanTypeID)
}

func (r *LoanRepository) ListLoansByUser(ctx context.Context, userID pgtype.Int4) ([]sqlc.Loan, error) {
	return r.queries.ListLoansByUser(ctx, userID)
}

func (r *LoanRepository) UpdateLoan(ctx context.Context, loan sqlc.UpdateLoanParams) error {
	return r.queries.UpdateLoan(ctx, loan)
}

func (r *LoanRepository) DeleteLoan(ctx context.Context, id int32) error {
	return r.queries.DeleteLoan(ctx, id)
}

func (r *LoanRepository) DeleteLoansByAccount(ctx context.Context, accountID pgtype.Int4) error {
	return r.queries.DeleteLoansByAccount(ctx, accountID)
}

func (r *LoanRepository) DeleteLoansByType(ctx context.Context, loanTypeID pgtype.Int4) error {
	return r.queries.DeleteLoansByType(ctx, loanTypeID)
}

func (r *LoanRepository) DeleteLoansByUser(ctx context.Context, userID pgtype.Int4) error {
	return r.queries.DeleteLoansByUser(ctx, userID)
}
