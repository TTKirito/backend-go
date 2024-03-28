migrate:
	migrate create -ext sql -dir db/migration -seq add_user
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
mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/TTKirito/backend-go/db/sqlc Store
.PHONY: migrate