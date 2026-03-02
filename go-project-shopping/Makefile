include .env
export

# Configurations
MIGRATION_DIR= internal/db/migrations

COMPOSE_FILE_PROD=docker-compose.prod.yaml
COMPOSE_FILE_DEV=docker-compose.dev.yaml
ENV_FILE=.env
ENV_FILE_EXAMPLE=.env.example

# ===================== App commands =====================
# Run Server
run-server:
	go run ./cmd/api

# Run worker
run-worker:
	go run ./cmd/worker

# Build Server
build:
	go build -o bin/build ./cmd/api

# Run built server
run-built:
	./bin/build

# Sqlc generate
sqlc:
	sqlc generate

# Create new migration: (make create-migrate NAME=users_table)
create-migrate:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(NAME)

# Apply migrations: (make migrate-up)
migrate-up:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE_URL)" up

# Rollback migrations: (make migrate-down)
migrate-down:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE_URL)" down 1

# Rollback N migrations: (make migrate-down-n N=2)
migrate-down-n:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE_URL)" down $(N)

# Force migration to specific version: (make migrate-force VERSION=2)
migrate-force:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE_URL)" force $(VERSION)

# Drop all tables
migrate-drop:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE_URL)" drop

# Migrate goto version: (make migrate-goto VERSION=2)
migrate-goto:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE_URL)" goto $(VERSION)


# ===================== Docker commands =====================
# Import 
importdb:
	docker exec -i postgres-db psql -U postgres -d gin < ./db/backup.sql

# Backup
backupdb:
	docker exec -i postgres-db pg_dump -U postgres -d gin > ./db/backup.sql

# ===================== Production commands ======================
# Build and run production server
prod:
	docker compose -f $(COMPOSE_FILE_PROD) down
	docker compose -f $(COMPOSE_FILE_PROD) --env-file $(ENV_FILE_EXAMPLE) up -d --build

# Stop production server
stop-prod:
	docker compose -f $(COMPOSE_FILE_PROD) down

# View logs (last 100 lines and follow)
logs-prod:
	docker compose -f $(COMPOSE_FILE_PROD) logs -f --tail=100

# Access bash in running container
bash:
	docker exec -it go-app sh

# ===================== Dev commands =====================
# Build and run dev server
dev:
	docker compose -f $(COMPOSE_FILE_DEV) down
	docker compose -f $(COMPOSE_FILE_DEV) --env-file $(ENV_FILE) up -d --build

# Stop dev server
stop-dev:
	docker compose -f $(COMPOSE_FILE_DEV) down

.PHONY: importdb backupdb server create-migrate migrate-up migrate-down migrate-down-n migrate-force migrate-drop migrate-goto sqlc run build run-built prod stop-prod logs-prod bash dev stop-dev