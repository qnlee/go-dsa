.PHONY: test build coverage

build:
	go build -v ./...

test_leetcode_medium:
	go test -v ./leetcode/medium/...

bench_leetcode_medium:
	go test -bench=. ./leetcode/medium/... -run=^$ -v

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

bench:
	go test -bench=. ./... -run=^$ -v
