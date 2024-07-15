ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build:
	@go build -o bin/waterwatch_back ./cmd/waterwatch_back

run: build
	@./bin/waterwatch_back

deploy: docker-up migrate-up run

test:
	@go test -v ./...

# Docker
docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

# Goose
migrate-up:
	@cd migration && goose postgres "host=localhost user=$$DB_USER dbname=$$DB_NAME password=$$DB_PASSWORD sslmode=disable" up && cd ..

migrate-down:
	@cd migration && goose postgres "host=localhost user=$$DB_USER dbname=$$DB_NAME password=$$DB_PASSWORD sslmode=disable" down-to 0 && cd ..
