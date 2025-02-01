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
	"net/http"
	"os"

	server "github.com/nikashlabs/hishab/internal/server"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello bro this is awesome")
}

func main() {
	databaseURL, exists := os.LookupEnv("DATABASE_URL")

	if !exists {
		fmt.Fprintf(os.Stderr, "DATABASE_URL not set\n")
		os.Exit(1)
	}

	databaseConnection, err := server.ConnectDatabase(databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up database: %v\n", err)
		os.Exit(1)
	}
	defer databaseConnection.Close(context.Background())

	fmt.Println("Database connected successfully")

	router := http.NewServeMux()
	router.HandleFunc("GET /", rootHandler)

	http.ListenAndServe(":5000", router)
}
