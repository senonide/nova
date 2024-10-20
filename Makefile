lint:
	@golangci-lint run

test:
	@go test -race -count=1 ./...

build-debug:
	@go build -gcflags="all=-N -l" -o nova-debug

build:
	@go build -gcflags="all=-N -l" -o nova

run:
	@go run .
