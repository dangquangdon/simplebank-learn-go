migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@127.0.0.1:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@127.0.0.1:5432/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

fmt-all:
	go fmt ./...

.PHONY: migrateup migratedown sqlc