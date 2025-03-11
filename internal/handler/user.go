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
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nikashlabs/hishab/internal/database/sqlc"
	"github.com/nikashlabs/hishab/internal/repository"
	"github.com/nikashlabs/hishab/pkg/logger"
)

// sample: how to use repository inside handler functions to interact with database

func HandleUser(logger logger.Logger, userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUserByID(w, r, logger, userRepo)
		case http.MethodPost:
			createUser(w, r, logger, userRepo)
		case http.MethodPut:
			updateUser(w, r, logger, userRepo)
		case http.MethodDelete:
			deleteUser(w, r, logger, userRepo)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request, logger logger.Logger, userRepo *repository.UserRepository) {
	var user sqlc.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("failed to decode request body", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdUser, err := userRepo.CreateUser(r.Context(), user)
	if err != nil {
		logger.Error("failed to create user", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func getUserByID(w http.ResponseWriter, r *http.Request, logger logger.Logger, userRepo *repository.UserRepository) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error("invalid user id", "error", err)
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := userRepo.GetUserByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to fetch user", "error", err)
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request, logger logger.Logger, userRepo *repository.UserRepository) {
	var user sqlc.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("failed to decode request body", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := userRepo.UpdateUser(r.Context(), user); err != nil {
		logger.Error("failed to update user", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user updated successfully"})
}

func deleteUser(w http.ResponseWriter, r *http.Request, logger logger.Logger, userRepo *repository.UserRepository) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error("invalid user id", "error", err)
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if err := userRepo.DeleteUser(r.Context(), id); err != nil {
		logger.Error("failed to delete user", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "user deleted successfully"})
}
