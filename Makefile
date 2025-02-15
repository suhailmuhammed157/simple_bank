postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_DB=simple_bank -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createDb:
	docker exec -it postgres12 createdb -U root -O root simple_bank

dropDb:
	docker exec -it postgres12 dropdb simple_bank

migrateUp:
	migrate -path database/migrations -database "postgres://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateDown:
	migrate -path database/migrations -database "postgres://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createDb postgres dropDb migrateUp migrateDown sqlc test