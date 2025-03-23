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
	"github.com/redis/go-redis/v9"

	"github.com/nikashlabs/hishab/internal/cache"
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
			log.Error("missing required environment variable", "variable", key)
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
		return fmt.Errorf("failed to load server configuration, %w", err)
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
			log.Fatal("error listening and serving", "reason", err)
		}
	}()

	return handleGracefulShutdown(log, httpServer, ctx)
}

func handleGracefulShutdown(log logger.Logger, httpServer *http.Server, ctx context.Context) error {
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Error("failed to shutdown server", "reason", err)
		return err
	}

	log.Info("server shutdown gracefully")
	return nil
}

// change here: remove it when redis implementation is done
func testRedis(ctx context.Context, log logger.Logger, rdb *redis.Client) {
	key := "greatness"
	value := "nothingness"

	if err := rdb.Set(ctx, key, value, 0).Err(); err != nil {
		log.Error("failed to set key in redis", "reason", err)
		return
	}
	log.Info("Set value in cache", key, value)

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Error("failed to get key from redis", "reason", err)
	}
	log.Info("Got value from cache", key, val)
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
	start := time.Now() // Start timing

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// logger initialization
	// change here: if you want to use a different logger
	log, err := logger.NewZapLogger()
	if err != nil {
		return fmt.Errorf("failed to initialize logger, %w", err)
	}

	if flushable, ok := log.(logger.Flushable); ok {
		defer func() {
			if err := flushable.Sync(); err != nil {
				fmt.Fprintf(os.Stderr, "failed to flush logs, %v\n", err)
			}
		}() // ensure flush, for flushable loggers
	}

	/*
		// load .env (not needed, when using docker-compose)
		if err := godotenv.Load(); err != nil {
			return fmt.Errorf("failed to load .env, %w", err)
		}
	*/

	// database initialization
	databaseConnectionPool, err := setupDatabase(log)
	if err != nil {
		return err
	}
	defer databaseConnectionPool.Close()

	// cache initialization
	rdb, err := cache.Init(log)
	if err != nil {
		log.Fatal("failed to initialize redis")
		return err
	}
	defer rdb.Close()

	// change here: remove it when redis implementation is done
	testRedis(ctx, log, rdb)

	// repositories initialization
	repos := setupRepositories(databaseConnectionPool)

	// time taken for prerequisites (logger, db, cache) loading
	elapsed := time.Since(start)
	log.Info("server prerequisites (logger, db, cache) loaded successfully", "time", elapsed.String())

	return setupServer(log, ctx, repos)
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
