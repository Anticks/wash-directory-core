# Define default variables (can be overridden when running `make`)
MIGRATION_NAME ?= initial

# Command to generate a new migration with a dynamic name
migrate-new:
	@if [ -z "$(name)" ]; then \
		echo "Please provide a migration name using 'make migrate-new name=<migration_name>'"; \
		exit 1; \
	fi; \
	atlas migrate diff $(name) \
	  --dir "file://migrations" \
	  --to "ent://ent/schema" \
	  --dev-url "postgres://postgres:password@localhost:5432/directory_core?sslmode=disable&search_path=public"

# Apply the migrations to the database
migrate-apply:
	atlas migrate apply \
	  --dir "file://migrations" \
	  --url "postgres://postgres:password@localhost:5432/directory_core?sslmode=disable&search_path=public"

# Check migration status
migrate-status:
	atlas migrate status \
	  --dir "file://migrations" \
	  --url "postgres://postgres:password@localhost:5432/directory_core?sslmode=disable&search_path=public"

# Clean up migration files (for development purposes)
migrate-clean:
	rm -rf migrations/*

# Generate Ent schema
generate-ent:
	go run entgo.io/ent/cmd/ent generate ./ent/schema

.PHONY: dev air templ tailwind clean build run

dev: air templ tailwind
	@echo "All development tasks started in parallel. Press Ctrl+C to stop."

# Run Air for live reloading in parallel
air:
	@echo "Starting Air (Go live reload)..."
	@air &

# Watch for changes in Templ files in parallel
templ:
	@echo "Starting Templ watch..."
	@templ watch &

# Watch for Tailwind CSS changes in parallel
tailwind:
	@echo "Starting Tailwind CSS watch..."
	@npx tailwindcss -i ./static/css/tailwind.css -o ./static/css/output.css --watch &

# Cleanup any temporary files (optional)
clean:
	@echo "Cleaning project..."
	@rm -rf ./build

# Build the project (optional)
build:
	@echo "Building project..."
	@go build -o app ./cmd/main.go

# Run the project (optional)
run:
	@echo "Running project..."
	@./app

