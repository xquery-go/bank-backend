postgres:
	docker run --name latest_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it latest_postgres createdb --username=root --owner=root go_banking

dropdb:
	docker exec -it latest_postgres dropdb go_banking

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_banking?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_banking?sslmode=disable" -verbose down

.PHONY: createdb postgres dropdb migrateup migratedown


