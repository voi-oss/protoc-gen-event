# protoc-gen-event

`protoc-gen-event` is a [`protoc`](https://grpc.io/docs/protoc-installation/) plugin to make publishing *proto-defined*
events easier.

* events are described with proto files
* Go code representing the events is generated
* Watermill publication and consumption functions are generated

## How to use

### Installation

Download `protoc-gen-event` and make sure it is available in your PATH:

```
go install github.com/voi-oss/protoc-gen-event/cmd/protoc-gen-event@latest
```

### Usage

Check the [`./examples/`](./examples) folder to see how events are defined and generate the associated Go code with:

```
protoc -I . -I /usr/include \
    --event_out=. \
    --event_opt=suffixMatch=Event,paths=source_relative \
    --go_out=paths=source_relative:. \
    ./examples/events.proto
```

### Parameters

`protoc --event_out=. --event_opt=sufficMatch=EventMatchSuffix,paths=source_relative`

## Development

### Building the event options type

```
go build -o protoc-gen-event ./cmd/protoc-gen-event/* &&
protoc -I . -I /usr/include \
    --plugin protoc-gen-event=./protoc-gen-event \
    --go_out=paths=source_relative:. \
    ./pkg/options/descriptor.proto
```

### Running protoc-gen-event

```
go build -o protoc-gen-event ./cmd/protoc-gen-event/* &&
protoc -I . -I /usr/include \
    --plugin protoc-gen-event=./protoc-gen-event \
    --event_out=. \
    --event_opt=suffixMatch=Event,requiredFields=string:messageID+google.protobuf.Timestamp:generatedAt,paths=source_relative \
    --go_out=paths=source_relative:. \
    ./examples/events.proto
```

## Contributions

We encourage and support an active, healthy community of contributors &mdash;
including you! Details are in the [contribution guide](CONTRIBUTING.md) and
the [code of conduct](CODE_OF_CONDUCT.md). The maintainers keep an eye on
issues and pull requests, but you can also report any negative conduct to
opensource@voiapp.io.

### Contributors

- [@K-Phoen](https://github.com/K-Phoen)
- [@drpytho](https://github.com/drpytho)

## License

Apache 2.0, see [LICENSE](LICENSE).