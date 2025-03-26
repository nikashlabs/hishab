.PHONY: dev dev-build dev-detached dev-build-detached prod prod-build prod-detached prod-build-detached stop stop-detached lint format clean reset reset-dev help

DEV_COMPOSE_FILES := -f docker-compose.dev.yaml

# Default target showing help
help:
	@echo "Commands:"
	@echo "Development:"
	@echo "  make dev                 		- Run application in development mode"
	@echo "  make dev-build           		- Rebuild and run development environment"
	@echo "  make dev-detached        		- Run development environment in detached mode"
	@echo "  make dev-build-detached  		- Rebuild and run development environment in detached mode"
	@echo "  make dev-build-detached  		- Rebuild and run development environment in detached mode"
	@echo ""
	@echo "Production:"
	@echo "  make prod                		- Run application in production mode"
	@echo "  make prod-build          		- Rebuild and run production environment"
	@echo "  make prod-detached       		- Run production environment in detached mode"
	@echo "  make prod-build-detached 		- Rebuild and run production environment in detached mode"
	@echo ""
	@echo "Containers:"
	@echo "  make stop                		- Stop and remove all containers"
	@echo "  make reset       				- Stops all containers, removes volumes and orphans, and rebuilds the Docker environment"
	@echo "  make reset-dev      			- Reset, for dev containers"
	@echo ""
	@echo "Tools:"
	@echo "  make lint                		- Run linters on Go code"
	@echo "  make format              		- Format Go code using gofmt"
	@echo "  make clean               		- Remove build artifacts"

# Detect OS
ifeq ($(OS),Windows_NT)
  DETECTED_OS := Windows
else
  DETECTED_OS := $(shell uname -s)
endif

# Development environment
dev:
ifeq ($(DETECTED_OS),Windows)
	docker compose $(DEV_COMPOSE_FILES) up
else
	COMPOSE_BAKE=true docker compose $(DEV_COMPOSE_FILES) up
endif

dev-build:
ifeq ($(DETECTED_OS),Windows)
	docker compose $(DEV_COMPOSE_FILES) up --build
else
	COMPOSE_BAKE=true docker compose $(DEV_COMPOSE_FILES) up --build
endif

dev-detached:
ifeq ($(DETECTED_OS),Windows)
	docker compose $(DEV_COMPOSE_FILES) up -d
else
	COMPOSE_BAKE=true docker compose $(DEV_COMPOSE_FILES) up -d
endif

dev-build-detached:
ifeq ($(DETECTED_OS),Windows)
	docker compose $(DEV_COMPOSE_FILES) up -d --build
else
	COMPOSE_BAKE=true docker compose $(DEV_COMPOSE_FILES) up -d --build
endif

# Production environment
prod:
	docker compose up

prod-build:
	docker compose up --build

prod-detached:
	docker compose up -d

prod-build-detached:
	docker compose up -d --build

# Container Management
stop:
	docker compose down

reset:
	docker compose down -v --remove-orphans && docker compose up --build

reset-dev:
ifeq ($(DETECTED_OS),Windows)
	docker compose down -v --remove-orphans && docker compose $(DEV_COMPOSE_FILES) up --build
else
	COMPOSE_BAKE=true docker compose down -v --remove-orphans && docker compose $(DEV_COMPOSE_FILES) up --build
endif

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
