GO_BIN := $(shell go env GOPATH)
export PATH := ${GO_BIN}/bin:$(PATH)

new-migration:
	go install github.com/pressly/goose/v3/cmd/goose@latest &&	goose -dir "./migrations" create new_migration sql

migrate-up:
	go install github.com/pressly/goose/v3/cmd/goose@latest && goose --dir "./migrations" postgres "user=root dbname=shooters host=localhost password=root port=5445 sslmode=disable" up

migrate-down:
	go install github.com/pressly/goose/v3/cmd/goose@latest && goose --dir "./migrations" postgres "user=root dbname=shooters host=localhost password=root port=5445 sslmode=disable" down
