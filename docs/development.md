# Development

## Building the event options type

```
go build -o protoc-gen-event ./cmd/protoc-gen-event/* &&
protoc -I . -I /usr/include \
    --plugin protoc-gen-event=./protoc-gen-event \
    --go_out=paths=source_relative:. \
    ./pkg/options/descriptor.proto
```

## Running protoc-gen-event

```
go build -o protoc-gen-event ./cmd/protoc-gen-event/* &&
protoc -I . -I /usr/include \
    --plugin protoc-gen-event=./protoc-gen-event \
    --event_out=. \
    --event_opt=suffixMatch=Event,requiredFields=string:messageID+google.protobuf.Timestamp:generatedAt,paths=source_relative \
    --go_out=paths=source_relative:. \
    ./examples/events.proto
```