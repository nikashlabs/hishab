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
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nikashlabs/hishab/internal/database"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello bro this is awesome")
}

func loadConfig() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env\n")
	}
}

func setupDatabase() {
	// Initialize database
	status, databaseConnectionPool := database.Init()
	if status {
		defer databaseConnectionPool.Close()
	}
}

func setupServer() {
	// Initialize server
	router := http.NewServeMux()
	router.HandleFunc("/", rootHandler)
	http.ListenAndServe(":5000", router)
}

func main() {
	loadConfig()
	setupDatabase()
	setupServer()
}
