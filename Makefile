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

# run tests excluding some files (sqlc generated)
COVERAGE_EXCLUDE := "db.go\|models.go"
test:
	go test -v -coverprofile=coverage.tmp ./... && \
		cat coverage.tmp | grep -v $(COVERAGE_EXCLUDE) > coverage.out && \
		go tool cover -func=coverage.out && \
		rm coverage.tmp coverage.out

test-html:
	go test -v -coverprofile=coverage.tmp ./... && \
		cat coverage.tmp | grep -v $(COVERAGE_EXCLUDE) > coverage.out && \
		go tool cover -html=coverage.out && \
		rm coverage.tmp coverage.out

.PHONY: createdb postgres dropdb migrateup migratedown sqlc test test-html
