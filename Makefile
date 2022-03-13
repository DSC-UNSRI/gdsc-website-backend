
include .env
export $(shell sed 's/=.*//' .env)

run:
	go run cmd/main.go

build:
	go build -o dist/app cmd/main.go

tidy:
	go mod tidy

migrate-up:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_URL) -verbose up

migrate-down:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_URL) -verbose down 

migrate-fresh: migrate-down migrate-up

generate-migration-file:
	migrate create -ext sql -dir ./internal/db/postgresql/migration -seq $(table_name)

sqlc:
	sqlc generate

.PHONY:	migrate-up migrate-down sqlc migrate-fresh generate-migration-file run build