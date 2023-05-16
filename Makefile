.PHONY: test build coverage

build:
	go build -v ./...

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out