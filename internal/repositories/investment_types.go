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

type InvestmentTypeRepository struct {
	queries *sqlc.Queries
}

func NewInvestmentTypeRepository(db *pgxpool.Pool) *InvestmentTypeRepository {
	return &InvestmentTypeRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *InvestmentTypeRepository) CreateInvestmentType(ctx context.Context, investmentType sqlc.CreateInvestmentTypeParams) (sqlc.InvestmentType, error) {
	return r.queries.CreateInvestmentType(ctx, investmentType)
}

func (r *InvestmentTypeRepository) GetInvestmentTypeByID(ctx context.Context, id int32) (sqlc.InvestmentType, error) {
	return r.queries.GetInvestmentTypeByID(ctx, id)
}

func (r *InvestmentTypeRepository) GetInvestmentTypesByUserID(ctx context.Context, userID int32) ([]sqlc.InvestmentType, error) {
	return r.queries.GetInvestmentTypesByUserID(ctx, userID)
}

func (r *InvestmentTypeRepository) UpdateInvestmentType(ctx context.Context, investmentType sqlc.UpdateInvestmentTypeParams) error {
	return r.queries.UpdateInvestmentType(ctx, investmentType)
}

func (r *InvestmentTypeRepository) DeleteInvestmentType(ctx context.Context, id int32) error {
	return r.queries.DeleteInvestmentType(ctx, id)
}
