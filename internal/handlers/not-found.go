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
package handlers

import (
	"net/http"

	"github.com/nikashlabs/hishab/pkg/logger"
)

// Template for creating a new handler
// func handleSomething(logger logger.Logger) http.Handler {
// 	thing := prepareThing()
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			// use thing to handle request
// 			logger.Info(r.Context(), "msg", "handleSomething")
// 		}
// 	)
// }

func HandleNotFound(logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wrong Path Bro\n"))
		logger.Info("Invalid Path", "invalid_path", r.URL.Path)
	})
}
