include .env
export

run:
	go run cmd/main.go

build:
	go build -o bin/scaleurl cmd/main.go

db-push:
	atlas schema apply --url "$(DATABASE_URL)" --to "file://schema/schema.sql" --dev-url "postgres://postgres:12345678@localhost:5432/scaleurl"

db-inspect:
	atlas schema inspect --env local
