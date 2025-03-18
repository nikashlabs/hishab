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

package server

import (
	"net/http"

	"github.com/nikashlabs/hishab/internal/repositories"
	"github.com/nikashlabs/hishab/pkg/logger"
)

func NewServer(
	logger logger.Logger,
	config *Config,
	repositories *repositories.Repositories,
) http.Handler {
	mux := http.NewServeMux()

	// Register routes
	addRoutes(
		mux,
		logger,
		config,
		repositories,
	)

	var httpHandler http.Handler = mux
	// httpHandler = someMiddleware(httpHandler)
	// httpHandler = someMiddleware2(httpHandler)
	// httpHandler = someMiddleware3(httpHandler)
	return httpHandler
}
