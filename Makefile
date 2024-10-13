build:
	@go build -o bin/go-cms

run: build
	@./bin/go-cms

test:
	@go test -v ./...