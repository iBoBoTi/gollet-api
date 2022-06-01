# create migration file
#	migrate create -ext sql -dir database/migrations -seq create_roles_table

include .env
export

run:
	go run cmd/main.go

generate-mock:
	go generate ./...

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_DATABASE)

dropdb:
	docker exec -it postgres14 dropdb $(DB_DATABASE)

migrate-create:
	migrate create -ext sql -dir database/migrations -seq walletApi_init_schema

migrate-up:
	migrate -path database/migrations -database $(DATABASE_URL) -verbose up

migrate-down:
	migrate -path database/migrations -database $(DATABASE_URL) -verbose down

migrate-clean:
	migrate -path database/migrations -database $(DATABASE_URL) force 1