-include .env

.PHONY: run
run:
	go run cmd/main.go -env=.env -config=config.yml

.PHONY: build
build:
	go build -o orgApi cmd/main.go -env=.env -config=config.yml

.PHONY: migrate-up
migrate-up:
	goose -dir ./migrations postgres "postgres://$(PG_USER):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_NAME)?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	goose -dir ./migrations postgres "postgres://$(PG_USER):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_NAME)?sslmode=disable" down
