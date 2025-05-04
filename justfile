set windows-shell := ["C:\\Program Files\\Git\\bin\\sh.exe", "-c"]

alias fd := frontend-dev
alias bd := backend-dev
alias i := install

default:
    # List all available commands
    just --list

frontend-install:
    # Install frontend dependencies
    cd frontend && bun install

backend-install:
    # Install backend dependencies
    cd backend/api/ && go install

frontend-lint:
    # Lint frontend
    cd frontend
    bun lint:fix

backend-lint:
    # Lint backend
    cd backend
    go mod tidy

frontend-dev:
    # Run frontend development server
    cd frontend && bunx astro dev

backend-dev:
    # Run backend development server
    cd backend/api/ && go run api.go

db-up:
    # Start local database in Docker
    cd backend
    bun run db:up

db-down:
    # Stop local database in Docker
    cd backend 
    bun run db:down

db-gen:
    # Generate SQL for database schema
    cd backend/cmd/server && go generate

# db-migrate:
#     # Apply database migrations
#     cd backend
#     bun migrate
# db-generate-migrate: db-generate db-migrate
#     # Apply SQL generation and migration
#     echo "Database schema generation and migration complete!"

backend-format:
    cd backend && go mod tidy

frontend-format:
    cd frontend && bun lint

install: frontend-install backend-install
    # Install frontend and backend dependencies
    echo "Frontend and backend dependencies installed!"

test:
    # Run tests in backend
    cd backend
    bun run test

generate:
    # Generate Schema from OpenAPI Specification
    cd backend && go generate
