ifeq ($(APP_MODE), development)
	-include .env
else
	include .env
endif

export $(shell sed 's/=.*//' .env)

run:
	go run cmd/main.go

build:
	CGO_ENABLED=0 go build -o dist/app cmd/main.go

tidy:
	go mod tidy

migrate-up:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_URL) -verbose up

migrate-down:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_URL) -verbose down 

migrate-drop:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_URL) -verbose drop

migrate-up-test:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_TEST_URL) -verbose up

migrate-down-test:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_TEST_URL) -verbose down -all

migrate-drop-test:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_TEST_URL) -verbose drop

migrate-fresh: migrate-down migrate-up

generate-migration-file:
	migrate create -ext sql -dir ./internal/db/postgresql/migration -seq $(table_name)

sqlc:
	sqlc -f sqlc.yaml generate

.PHONY:	migrate-up migrate-down sqlc migrate-fresh generate-migration-file run build