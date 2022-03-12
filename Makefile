
include .env
export $(shell sed 's/=.*//' .env)

run:
	go run cmd/main.go

tidy:
	go mod tidy

migrate-up:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_URL) -verbose up

migrate-down:
	migrate -path ./internal/db/postgresql/migration -database $(DB_CONNECTION_URL) -verbose down 

migrate-fresh: migrate-down migrate-up

generate-migration-file:
	migrate create -ext sql -dir ./db/migration -seq schema

sqlc:
	sqlc generate

.PHONY:	migrate-up migrate-down sqlc migrate-fresh generate-migration-file