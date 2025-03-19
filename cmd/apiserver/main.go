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

package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/nikashlabs/hishab/internal/database"
	"github.com/nikashlabs/hishab/internal/repositories"
	"github.com/nikashlabs/hishab/internal/server"
	"github.com/nikashlabs/hishab/pkg/logger"
)

func setupDatabase(log logger.Logger) (*pgxpool.Pool, error) {
	status, databaseConnectionPool := database.Init(log)
	if !status {
		return nil, fmt.Errorf("database initialization failed")
	}
	return databaseConnectionPool, nil
}

func loadServerConfig(log logger.Logger) (*server.Config, error) {
	// change here: while adding new config variables
	requiredVariables := []string{"HOST", "PORT"}
	variables := make(map[string]string)
	missingRequiredVariables := []string{}
	for _, key := range requiredVariables {
		value, exists := os.LookupEnv(key)
		if !exists || value == "" {
			log.Error("missing required environment variable: %s", key)
			missingRequiredVariables = append(missingRequiredVariables, key)
		}
		variables[key] = value
	}

	if len(missingRequiredVariables) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missingRequiredVariables)
	}

	// change here: while adding new config variables
	return &server.Config{
		Host: variables["HOST"],
		Port: variables["PORT"],
	}, nil
}

func setupServer(log logger.Logger, ctx context.Context, repos *repositories.Repositories) error {
	if repos == nil {
		return fmt.Errorf("repos cannot be nil")
	}

	// load config
	config, err := loadServerConfig(log)
	if err != nil {
		return fmt.Errorf("failed to load server configuration: %w", err)
	}

	// create
	srv := server.NewServer(log, config, repos)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	// start
	go func() {
		log.Info("server started", "address", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("error listening and serving: %s\n", err)
		}
	}()

	return handleGracefulShutdown(log, httpServer, ctx)
}

func handleGracefulShutdown(log logger.Logger, httpServer *http.Server, ctx context.Context) error {
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Error("failed to shutdown server", "error", err)
		return err
	}

	log.Info("server shutdown gracefully")
	return nil
}

func setupRepositories(databaseConnectionPool *pgxpool.Pool) *repositories.Repositories {
	return &repositories.Repositories{
		User:            repositories.NewUserRepository(databaseConnectionPool),
		Account:         repositories.NewAccountRepository(databaseConnectionPool),
		Currency:        repositories.NewCurrencyRepository(databaseConnectionPool),
		ExpenseCategory: repositories.NewExpenseCategoryRepository(databaseConnectionPool),
		ExpenseRecord:   repositories.NewExpenseRecordRepository(databaseConnectionPool),
		IncomeCategory:  repositories.NewIncomeCategoryRepository(databaseConnectionPool),
		IncomeRecord:    repositories.NewIncomeRecordRepository(databaseConnectionPool),
		Installment:     repositories.NewInstallmentRepository(databaseConnectionPool),
		InvestmentType:  repositories.NewInvestmentTypeRepository(databaseConnectionPool),
		Investment:      repositories.NewInvestmentRepository(databaseConnectionPool),
		LoanType:        repositories.NewLoanTypeRepository(databaseConnectionPool),
		Loan:            repositories.NewLoanRepository(databaseConnectionPool),
		ScheduledRecord: repositories.NewScheduledRecordRepository(databaseConnectionPool),
	}
}

func run(ctx context.Context, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// logger initialization
	// change here: if you want to use a different logger
	log, err := logger.NewZapLogger()
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	if flushable, ok := log.(logger.Flushable); ok {
		defer func() {
			if err := flushable.Sync(); err != nil {
				fmt.Fprintf(os.Stderr, "failed to flush logs: %v\n", err)
			}
		}() // ensure flush, for flushable loggers
	}

	// load .env
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env: %w", err)
	}

	// database initialization
	databaseConnectionPool, err := setupDatabase(log)
	if err != nil {
		return err
	}
	defer databaseConnectionPool.Close()

	// repositories initialization
	repos := setupRepositories(databaseConnectionPool)

	return setupServer(log, ctx, repos)
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
