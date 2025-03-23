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

package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nikashlabs/hishab/pkg/logger"
	"github.com/nikashlabs/hishab/pkg/util"
	"github.com/redis/go-redis/v9"
)

func Init(log logger.Logger) (*redis.Client, error) {
	// Load config
	options, err := loadRedisOptions(log)
	if err != nil {
		return nil, err
	}

	// Create client
	rdb := redis.NewClient(options)

	// Connect
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error("failed to connect to redis", "reason", err)
		return nil, err
	}
	log.Info("redis connection successful", "response", pong)
	return rdb, nil
}

func loadRedisOptions(log logger.Logger) (*redis.Options, error) {
	requiredVariables := []string{"REDIS_HOST", "REDIS_PORT", "REDIS_PASSWORD", "REDIS_DB", "REDIS_PROTOCOL"}

	variables := make(map[string]string)
	for _, key := range requiredVariables {
		value, exists := os.LookupEnv(key)
		if !exists || value == "" {
			log.Error("missing required environment variable", "variable", key)
			return nil, fmt.Errorf("missing required environment variable: %s", key)
		}
		variables[key] = value
	}

	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", variables["REDIS_HOST"], variables["REDIS_PORT"]),
		Password: variables["REDIS_PASSWORD"],
		DB:       util.Atoi(variables["REDIS_DB"]),
		Protocol: util.Atoi(variables["REDIS_PROTOCOL"]),
	}

	log.Info("redis options loaded", "addr", options.Addr, "db", options.DB, "protocol", options.Protocol)
	return options, nil
}
