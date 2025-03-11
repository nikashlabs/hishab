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

	"github.com/nikashlabs/hishab/internal/handler"
	"github.com/nikashlabs/hishab/internal/middlewares"
	"github.com/nikashlabs/hishab/internal/repository"
	"github.com/nikashlabs/hishab/pkg/logger"
)

func addRoutes(
	mux *http.ServeMux,
	logger logger.Logger,
	config *Config,
	repositories *repository.Repositories,
) {
	mux.Handle("/", middlewares.Log(logger, handler.HandleNotFound(logger)))

	// sample: how to pass repository to a handler
	mux.Handle("/user", middlewares.Log(logger, handler.HandleUser(logger, repositories.User)))

	// mux.Handle("/api/v1/comments", handleComments(logger, commentStore))
	// mux.Handle("/api/v1/another", handleAnother(logger, anotherStore))
	// mux.HandleFunc("/healthz", handleHealthz(logger))
	// mux.Handle("/", http.NotFoundHandler())
}
