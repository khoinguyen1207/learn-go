importdb:
	docker exec -i postgres-db psql -U postgres -d gin < ./db/backup.sql
backupdb:
	docker exec -i postgres-db pg_dump -U postgres -d gin > ./db/backup.sql
server:
	go run .