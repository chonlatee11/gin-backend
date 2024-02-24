server:
	go run main.go
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root $(name)
db-create-migration:
	migrate create -ext sql -dir db/migration/ $(name) -tz
db-migrate-up:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose up
db-migrate-down:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY: server db-create-migration db-migrate-up db-migrate-down postgres createdb sqlc