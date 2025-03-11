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

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikashlabs/hishab/internal/database/sqlc"
)

type LoanTypeRepository struct {
	queries *sqlc.Queries
}

func NewLoanTypeRepository(db *pgxpool.Pool) *LoanTypeRepository {
	return &LoanTypeRepository{
		queries: sqlc.New(db),
	}
}

func (r *LoanTypeRepository) CreateLoanType(ctx context.Context, loanType sqlc.CreateLoanTypeParams) (sqlc.LoanType, error) {
	return r.queries.CreateLoanType(ctx, loanType)
}

func (r *LoanTypeRepository) GetLoanTypeByID(ctx context.Context, id int32) (sqlc.LoanType, error) {
	return r.queries.GetLoanTypeByID(ctx, id)
}

func (r *LoanTypeRepository) GetLoanTypesByUserID(ctx context.Context, userID int32) ([]sqlc.LoanType, error) {
	return r.queries.GetLoanTypesByUserID(ctx, userID)
}

func (r *LoanTypeRepository) UpdateLoanType(ctx context.Context, loanType sqlc.UpdateLoanTypeParams) error {
	return r.queries.UpdateLoanType(ctx, loanType)
}

func (r *LoanTypeRepository) DeleteLoanType(ctx context.Context, id int32) error {
	return r.queries.DeleteLoanType(ctx, id)
}
