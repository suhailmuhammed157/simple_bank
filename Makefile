postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_DB=simple_bank -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createDb:
	docker exec -it postgres12 createdb -U root -O root simple_bank

dropDb:
	docker exec -it postgres12 dropdb simple_bank

migrateUp:
	migrate -path database/migrations -database "postgres://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up
	
migrateUp1:
	migrate -path database/migrations -database "postgres://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migrateDown:
	migrate -path database/migrations -database "postgres://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateDown1:
	migrate -path database/migrations -database "postgres://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

createMigration:
	migrate create -ext sql -dir ./database/migrations -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

redis:
	docker run -d --name redis -p 6379:6379 redis:alpine

proto:
	rm -f pb/*.go 
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: createDb postgres dropDb migrateUp migrateDown sqlc test server proto redis createMigration