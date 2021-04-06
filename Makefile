migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@127.0.0.1:5432/simplebank?sslmode=disable" -verbose up

migrateup-test:
	migrate -path db/migrations -database "postgresql://postgres:postgres@127.0.0.1:5433/simplebank_test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@127.0.0.1:5432/simplebank?sslmode=disable" -verbose down

migratedown-test:
	migrate -path db/migrations -database "postgresql://postgres:postgres@127.0.0.1:5433/simplebank_test?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./... -count=1

fmt-all:
	go fmt ./...

server:
	go run main.go

.PHONY: migrateup migratedown sqlc test fmt-all server migrateup-test migratedown-test