PHONY: lint
lint:
	golangci-lint run --config .golangci.yaml

PHONY: tests
tests:
	go test -v ./...

.PHONY: build
build:
	go build -o protoc-gen-event ./cmd/protoc-gen-event/*

PHONY: build_options
build_options: build
	protoc -I . -I /usr/include \
		--plugin protoc-gen-event=./protoc-gen-event \
		--go_out=paths=source_relative:. \
		./pkg/options/descriptor.proto

PHONY: build_examples
build_examples: build
	protoc -I . -I /usr/include \
		--plugin protoc-gen-event=./protoc-gen-event \
		--event_out=. \
		--event_opt=suffixMatch=Event,paths=source_relative \
		--go_out=paths=source_relative:. \
		./examples/events.proto