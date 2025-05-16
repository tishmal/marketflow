BINARY = marketflow

.PHONY: up down test fmt build run

up:
	docker compose -f deployments/docker-compose.yml up --build

down:
	docker compose -f deployments/docker-compose.yml down -v

test:
	go test ./...

fmt:
	go run mvdan.cc/gofumpt -w .

build:
	go build -o $(BINARY) .

run: build
	./$(BINARY) --help
