.PHONY: dev dev-build dev-detached dev-build-detached prod prod-build prod-detached prod-build-detached stop stop-detached lint format clean reset reset-dev help


DEV_COMPOSE_FILES := -f docker-compose.dev.yaml

# Default target showing help
help:
	@echo "Available commands:"
	@echo "Development commands:"
	@echo "  make dev                 - Run application in development mode"
	@echo "  make dev-build           - Rebuild and run development environment"
	@echo "  make dev-detached        - Run development environment in detached mode"
	@echo "  make dev-build-detached  - Rebuild and run development environment in detached mode"
	@echo ""
	@echo "Production commands:"
	@echo "  make prod                - Run application in production mode"
	@echo "  make prod-build          - Rebuild and run production environment"
	@echo "  make prod-detached       - Run production environment in detached mode"
	@echo "  make prod-build-detached - Rebuild and run production environment in detached mode"
	@echo ""
	@echo "Container management:"
	@echo "  make stop                - Stop and remove all containers"
	@echo "  make stop-detached       - Stop all containers without removing them"
	@echo ""
	@echo "Development tools:"
	@echo "  make lint                - Run linters on Go code"
	@echo "  make format              - Format Go code using gofmt"
	@echo "  make clean               - Remove build artifacts"

# Makefile Util

## Detect OS
ifeq ($(OS),Windows_NT)
  DETECTED_OS := Windows
  CMD_SEPARATOR := ;
else
  DETECTED_OS := $(shell uname -s)
  CMD_SEPARATOR := &&
endif

## Usage: $(call SET_ENV,VAR_NAME,VAR_VALUE)
define SET_ENV
  $(if $(filter Windows,$(DETECTED_OS)), \
    @powershell -Command "$$env:$(1)='$(2)';", \
    @export $(1)=$(2) \
  )
endef

## Usage: $(call PRINT_ENV,VAR_NAME)
define PRINT_ENV
  $(if $(filter Windows,$(DETECTED_OS)), \
    powershell -Command "Write-Output '$(1): $$env:$(1)'", \
    echo "$(1) is set to: $$$(1)" \
  )
endef

# Development environment
dev:
	$(call SET_ENV,COMPOSE_BAKE,true) $(CMD_SEPARATOR) docker compose $(DEV_COMPOSE_FILES) up

dev-build:
	$(call SET_ENV,COMPOSE_BAKE,true) $(CMD_SEPARATOR) docker compose $(DEV_COMPOSE_FILES) up --build

dev-detached:
	$(call SET_ENV,COMPOSE_BAKE,true) $(CMD_SEPARATOR) docker compose $(DEV_COMPOSE_FILES) up -d

dev-build-detached:
	$(call SET_ENV,COMPOSE_BAKE,true) $(CMD_SEPARATOR) docker compose $(DEV_COMPOSE_FILES) up -d --build

# Production environment
prod:
	docker compose up

prod-build:
	docker compose up --build

prod-detached:
	docker compose up -d

prod-build-detached:
	docker compose up -d --build

stop:
	docker compose down

reset:
	@docker compose down -v --remove-orphans $(CMD_SEPARATOR) docker compose up --build

reset-dev:
	$(call SET_ENV,COMPOSE_BAKE,true) $(CMD_SEPARATOR) docker compose down -v --remove-orphans $(CMD_SEPARATOR) docker compose $(DEV_COMPOSE_FILES) up --build

# Run linters - requires golangci-lint
lint:
	golangci-lint run

# Format code using gofmt
format:
	@echo "Formatting Go code..."
	gofmt -w -s .

# Clean build artifacts
clean:
	rm -rf ./tmp
