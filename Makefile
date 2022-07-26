PHONY: lint
lint:
	golangci-lint run --config .golangci.yaml

PHONY: tests
tests:
	go test -v ./...
