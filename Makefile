postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_DB=simple_bank -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createDb:
	docker exec -it postgres12 createdb -U root -O root simple_bank

dropDb:
	docker exec -it postgres12 dropdb simple_bank

migrateUp:
	migrate -path database/migrations -database "postgres://root:password@localhost:5433/simple_bank?sslmode=disable" -verbose up
	
migrateUp1:
	migrate -path database/migrations -database "postgres://root:password@localhost:5433/simple_bank?sslmode=disable" -verbose up 1

migrateDown:
	migrate -path database/migrations -database "postgres://root:password@localhost:5433/simple_bank?sslmode=disable" -verbose down

migrateDown1:
	migrate -path database/migrations -database "postgres://root:password@localhost:5433/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createDb postgres dropDb migrateUp migrateDown sqlc test