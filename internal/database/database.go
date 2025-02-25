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
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikashlabs/hishab/pkg/logger"
)

func Init() (bool, *pgxpool.Pool) {
	// Connection
	connectionPool, err := pgxpool.New(context.Background(), loadDatabaseURL())
	if err != nil {
		logger.Log.Error("Database connection failed: %v\n", err)
		return false, nil
	}
	logger.Log.Info("Database connection established")

	// Migrations
	err = RunMigrations(connectionPool)
	if err != nil {
		logger.Log.Error("Unable to run migrations: %v\n", err)
	} else {
		logger.Log.Info("Migrations ran successfully")
	}
	return true, connectionPool
}

func loadDatabaseURL() string {
	requiredVariables := []string{"POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_DB", "POSTGRES_SSLMODE"}

	variables := make(map[string]string)
	for _, key := range requiredVariables {
		value, exists := os.LookupEnv(key)
		if !exists || value == "" {
			logger.Log.Error("Missing required environment variable: %s", key)
			return ""
		}
		variables[key] = value
	}

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		variables["POSTGRES_USER"],
		variables["POSTGRES_PASSWORD"],
		variables["POSTGRES_HOST"],
		variables["POSTGRES_PORT"],
		variables["POSTGRES_DB"],
		variables["POSTGRES_SSLMODE"],
	)

	logger.Log.Info("Database URL:", databaseURL)
	return databaseURL
}

// func Connect(connectionString string) (*pgx.Conn, error) {
// 	conn, err := pgx.Connect(context.Background(), connectionString)
// 	if err != nil {
// 		return nil, fmt.Errorf("Unable to connect to database: %w", err)
// 	}
// 	return conn, nil
// }
