#!make
-include .env
export

default: up

bootstrap:
	go mod download
	go mod tidy
	if [ ! -f .env ]; then touch .env && cp .env.example .env; fi
	if [ ! -f ./sqlc/schema.sql ]; then cp sqlc/schema.sql.example sqlc/schema.sql; fi

migrate-lint:
	atlas migrate lint \
		--latest 1 \
		--dir "file://sqlc/migrations" \
		--dev-url "docker://postgres/12/test"

gen-migrate-up:
	atlas migrate diff $(NAME) \
		--dir "file://sqlc/migrations" \
		--to "file://sqlc/schema.sql" \
		--dev-url "docker://postgres/12/test"

migrate-up:
	atlas migrate apply $(N)\
		--url "postgres://postgres:${DB_PASSWORD}@localhost:5432/clean?sslmode=disable" \
		--dir "file://sqlc/migrations"

migrate-down:
	atlas migrate set $(V) \
		--url "postgres://postgres:${DB_PASSWORD}@localhost:5432/clean?sslmode=disable" \
		--dir "file://sqlc/migrations"

db-ui:
	atlas schema inspect \
		-u "postgres://postgres:${DB_PASSWORD}@localhost:5432/clean?sslmode=disable" \
		--web

up:
	docker compose -f compose.local.yaml up -d --force-recreate

up-hard:
	docker compose -f compose.local.yaml up -d --force-recreate --build

down:
	docker compose down

test:
	go test -v ./...

test-coverage:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out