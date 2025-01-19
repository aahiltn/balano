set shell := ["powershell.exe", "-c"]

frontend-install:
  # Install frontend dependencies
  cd frontend
  bun install

backend-install:
  # Install backend dependencies
  cd backend
  bun install

frontend-lint:
  # Lint frontend
  cd frontend
  bun lint:fix

backend-lint:
  # Lint backend
  cd backend
  bun lint:fix

frontend-dev:
  # Run frontend development server
  cd frontend 
  bun run start

backend-dev:
  # Run backend development server
  cd backend 
  bun run dev

backend:
  # Run backend production server with production database
  cd backend
  bun start

db-up:
  # Start local database in Docker
  cd backend
  bun run db:up

db-down:
  # Stop local database in Docker
  cd backend 
  bun run db:down

db-generate:
  # Generate SQL for database schema
  cd backend 
  bun generate

db-migrate:
  # Apply database migrations
  cd backend
  bun migrate

db-generate-migrate: db-generate db-migrate
  # Apply SQL generation and migration
  echo "Database schema generation and migration complete!"

format:
  # Format frontend and backend code
  cd backend
  bun format
  cd ../frontend
  bun format

install: frontend-install backend-install
  # Install frontend and backend dependencies
  echo "Frontend and backend dependencies installed!"

test:
  # Run tests in backend
  cd backend
  bun run test

generate:
  # Generate Schema from OpenAPI Specification
  cd backend
  bun run gen
