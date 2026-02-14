include .env
export

MIGRATION_DIR= internal/db/migrations

# Import 
importdb:
	docker exec -i postgres-db psql -U postgres -d gin < ./db/backup.sql

# Backup
backupdb:
	docker exec -i postgres-db pg_dump -U postgres -d gin > ./db/backup.sql

# Run Server
server:
	go run ./cmd/api/main.go

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

.PHONY: importdb backupdb server create-migrate migrate-up migrate-down migrate-down-n migrate-force migrate-drop migrate-goto sqlc