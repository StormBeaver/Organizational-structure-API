.PHONY: run
run:
	go run cmd/main.go

.PHONY: build
build:
	go build -o orgApi cmd/main.go

.PHONY: migrate-up
migrate-up:
	go run cmd/migrator/migrator.go

.PHONY: migrate-down
migrate-down:
	go run cmd/migrator/migrator.go -migration=false