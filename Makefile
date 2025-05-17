BINARY = marketflow

.PHONY: up down test fmt build run

include .env

up:
	docker compose --env-file .env -f deployments/docker-compose.yml up --build

down:
	docker compose --env-file .env -f deployments/docker-compose.yml down -v

test:
	go test ./...

fmt:
	go run mvdan.cc/gofumpt -w .

build:
	go build -o $(BINARY) .

run: build
	./$(BINARY) --help
