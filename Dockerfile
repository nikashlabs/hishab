# Copyright (C) 2025 Nikash Labs
# 
# This file is part of hishab.
# 
# hishab is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
# 
# hishab is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
# 
# You should have received a copy of the GNU General Public License
# along with hishab.  If not, see <https://www.gnu.org/licenses/>.

FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/apiserver ./cmd/apiserver
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/apiserver /app/apiserver
COPY --from=builder /app/.env /app/.env
COPY --from=builder /app/internal/database/migrations /app/internal/database/migrations
EXPOSE ${PORT}
CMD ["./apiserver"]
