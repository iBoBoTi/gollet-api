# create migration file
#	migrate create -ext sql -dir db/migrations -seq create_roles_table

include .env
export

run:
	go run cmd/main.go

generate-mock:
	go generate ./...

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_DATABASE)

dropdb:
	docker exec -it postgres14 dropdb $(POSTGRES_DATABASE)

migrate-up:
	migrate -path db/migrations -database $(DATABASE_URL) -verbose up

migrate-down:
	migrate -path db/migrations -database $(DATABASE_URL) -verbose down

migrate-clean:
	migrate -path db/migrations -database $(DATABASE_URL) force 1