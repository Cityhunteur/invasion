
build:
	go build ./...
.PHONY: build

lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint
	@golangci-lint run --config .golangci.yaml
.PHONY: lint

test:
	@go test -v \
		-count=1 \
		-timeout=5m \
		./...
.PHONY: test

test-profile:
	@go test \
		-shuffle=on \
		-count=1 \
		-race \
		-timeout=10m \
		./... \
		-coverprofile=coverage.out
.PHONY: test-profile

coverage:
	@go tool cover -func=./coverage.out
.PHONY: coverage

run:
	@echo Running simulation using default values...
	@go run main.go --aliens 10 --map testdata/example.map
.PHONY: run
