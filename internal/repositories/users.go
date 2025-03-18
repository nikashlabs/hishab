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

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: sqlc.New(db),
	}
}

// TODO: validate data before setting

func (r *UserRepository) CreateUser(ctx context.Context, user sqlc.User) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		IsVerified: user.IsVerified,
		PhotoUrl:   user.PhotoUrl,
	})
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (sqlc.User, error) {
	return r.queries.GetUserByID(ctx, int32(id))
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	return r.queries.GetUserByEmail(ctx, email)
}

func (r *UserRepository) UpdateUser(ctx context.Context, user sqlc.User) error {
	return r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		IsVerified: user.IsVerified,
		PhotoUrl:   user.PhotoUrl,
	})
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	return r.queries.DeleteUser(ctx, int32(id))
}

func (r *UserRepository) MarkUserAsVerified(ctx context.Context, id int) error {
	return r.queries.MarkUserAsVerified(ctx, int32(id))
}

func (r *UserRepository) UpdateLastActive(ctx context.Context, id int) error {
	return r.queries.UpdateLastActive(ctx, int32(id))
}
