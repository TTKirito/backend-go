migrate:
	migrate create -ext sql -dir db/migration -seq init_schema
migrateup:
	migrate -path db/migration -database "postgres://postgres:changeme@localhost:5434/postgres?sslmode=disable" --verbose up
migratedown:
	migrate -path db/migration -database "postgres://postgres:changeme@localhost:5434/postgres?sslmode=disable" --verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: migrate