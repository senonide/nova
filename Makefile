.PHONY: lint test build-debug build run server clean help

# Default target
help:
	@echo "Available commands:"
	@echo "  lint        - Run golangci-lint"
	@echo "  test        - Run tests with race detection"
	@echo "  build       - Build the server binary"
	@echo "  build-debug - Build debug version"
	@echo "  run         - Run the server directly with go run"
	@echo "  server      - Build and run the server binary"
	@echo "  clean       - Clean build artifacts"

lint:
	@golangci-lint run --config .golangci.yml

test:
	@go test -race -count=1 ./...

build-debug:
	@go build -gcflags="all=-N -l" -o nova-debug

build:
	@go build -o nova-server cmd/server/main.go

run:
	@go run cmd/server/main.go

server: build
	@echo "Starting Nova server on port 8080..."
	@./nova-server

clean:
	@rm -f nova-server nova-debug
	@echo "Cleaned build artifacts"
