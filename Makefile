# load env files
include .env
export

# create the container
postgres:
	docker run --name sbank_postgres \
		-p 5432:5432 \
		-e POSTGRES_USER=${POSTGRES_USER} \
		-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
		-d postgres:latest

# create db inside container
createdb:
	docker exec -it sbank_postgres createdb --username=${POSTGRES_USER} --owner=root ${POSTGRES_DB_NAME}

# drop db inside container
dropdb:
	docker exec -it sbank_postgres dropdb ${POSTGRES_DB_NAME}

# migrate db up
migrateup:
	migrate -path db/migration -database "${POSTGRES_URL}" -verbose up

# migrate db down
migratedown:
	migrate -path db/migration -database "${POSTGRES_URL}" -verbose down

# generate sqlc queries
sqlc:
	sqlc generate

# run tests
test:
	go test -v -cover ./...

.PHONY: createdb postgres dropdb migrateup migratedown sqlc test
