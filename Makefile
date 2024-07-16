ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Set Go version
GO_VERSION := 1.22.3

build: ensure-go-version
	@go build -o bin/waterwatch_back ./cmd/waterwatch_back

run: ensure-go-version build
	@./bin/waterwatch_back

deploy: ensure-go-version docker-up migrate-up run

test: ensure-go-version
	@go test -v ./...

# Docker
docker-up: ensure-go-version
	@docker-compose up -d

docker-down: ensure-go-version
	@docker-compose down

# Goose
GOOSE := $(shell which goose || (go install github.com/pressly/goose/v3/cmd/goose@latest && echo $(HOME)/go/bin/goose))

migrate-up: ensure-go-version
	@cd migration && $(GOOSE) postgres "host=localhost user=$$DB_USER dbname=$$DB_NAME password=$$DB_PASSWORD sslmode=disable" up && cd ..

migrate-down: ensure-go-version
	@cd migration && $(GOOSE) postgres "host=localhost user=$$DB_USER dbname=$$DB_NAME password=$$DB_PASSWORD sslmode=disable" down-to 0 && cd ..

migrate-restart: migrate-down migrate-up

# Ensure Go version
ensure-go-version:
	@go version | grep -q $(GO_VERSION) || (echo "Installing Go $(GO_VERSION)..." && \
	curl -LO https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz && \
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz && \
	rm go$(GO_VERSION).linux-amd64.tar.gz && \
	echo "export PATH=/usr/local/go/bin:\$$PATH" >> ~/.profile && \
	. ~/.profile && \
	go version)

# Run ensure-go-version before any target
.PHONY: build run deploy test docker-up docker-down migrate-up migrate-down migrate-restart ensure-go-version
