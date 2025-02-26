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

	"github.com/joho/godotenv"

	"github.com/nikashlabs/hishab/internal/database"
	"github.com/nikashlabs/hishab/internal/server"
	"github.com/nikashlabs/hishab/pkg/logger"
)

func loadEnv(log logger.Logger) {
	err := godotenv.Load()
	if err != nil {
		log.Error("Could not load .env file")
	}
}

func setupDatabase(log logger.Logger) {
	status, databaseConnectionPool := database.Init(log)
	if status {
		defer databaseConnectionPool.Close()
	}
}

func loadServerConfig(log logger.Logger) *server.Config {
	// change here: while adding new config variables
	requiredVariables := []string{"HOST", "PORT"}

	variables := make(map[string]string)
	for _, key := range requiredVariables {
		value, exists := os.LookupEnv(key)
		if !exists || value == "" {
			log.Error("Missing required environment variable: %s", key)
		}
		variables[key] = value
	}
	// change here: while adding new config variables
	return &server.Config{
		Host: variables["HOST"],
		Port: variables["PORT"],
	}
}

func setupServer(log logger.Logger, ctx context.Context) error {
	config := loadServerConfig(log)

	// Create
	srv := server.NewServer(log, config)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	// Start
	go func() {
		log.Info("Server started", "address", httpServer.Addr)
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
		log.Error("Could not shutdown HTTP server", "error", err)
		return err
	}

	log.Info("HTTP server shutdown gracefully")
	return nil
}

func run(ctx context.Context, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// change here: if you want to use a different logger
	log, err := logger.NewZapLogger()
	if err != nil {
		return err
	}

	if flushable, ok := log.(logger.Flushable); ok {
		defer flushable.Sync() // ensure flush, for flushable loggers
	}

	loadEnv(log)
	setupDatabase(log)
	return setupServer(log, ctx)
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
