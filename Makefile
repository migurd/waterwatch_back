build:
	@go build -o bin/waterwatch_back

run: build
	@./bin/waterwatch_back

test:
	@go test -v ./...