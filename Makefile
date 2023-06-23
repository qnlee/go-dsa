.PHONY: test build coverage

build:
	go build -v ./...
test_lc_easy:
	go test -v -cover ./leetcode/easy/...

test_lc_medium:
	go test -v -cover ./leetcode/medium/...

bench_lc_easy:
	go test -bench=. ./leetcode/easy/... -run=^$ -v

bench_lc_medium:
	go test -bench=. ./leetcode/medium/... -run=^$ -v

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

bench:
	go test -bench=. ./... -run=^$
