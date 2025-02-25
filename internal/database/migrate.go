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

package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/nikashlabs/hishab/pkg/logger"
)

func RunMigrations(databaseConnectionPool *pgxpool.Pool) error {
	// Convert *pgxpool.Pool to *sql.DB
	// https://github.com/jackc/pgx/blob/master/stdlib/sql.go
	db := stdlib.OpenDBFromPool(databaseConnectionPool)

	// Create migration driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("unable to create migration driver: %w", err)
	}

	// Create migration instance
	migration_instance, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("unable to create migration instance: %w", err)
	}

	// Apply migrations
	err = migration_instance.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("unable to apply migrations: %w", err)
	}

	logger.Log.Info("Migrations applied successfully!")
	return nil
}
