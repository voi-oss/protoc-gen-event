# protoc-gen-event

`protoc-gen-event` is a [`protoc`](https://grpc.io/docs/protoc-installation/) plugin to make publishing *proto-defined*
events easier. 

* events are described with proto files
* Go code representing the events is generated
* Watermill publication and consumption functions are generated 

## How to use

### Installation

`go get github.com/voi-oss/protoc-gen-event`

### Generating the code

Check the [`./examples/`](./examples) folder to see how events are defined and generate the associated Go code with:

```
go build -o protoc-gen-event ./cmd/protoc-gen-event/*
protoc --plugin ./protoc-gen-event -I . -I /usr/include --event_out=paths=source_relative:. --go_out=paths=source_relative:. ./examples/events.proto
```

### Parameters
`protoc --event_out=sufficMatch=EventMatchSuffix,paths=source_relative:.`

## Development

### Building the event options type

```
go build -o protoc-gen-event ./cmd/protoc-gen-event/*
protoc --plugin ./protoc-gen-event -I . -I /usr/include --go_out=paths=source_relative:. ./pkg/options/descriptor.proto
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