# Defining events

With `protoc-gen-event`, Watermill messages can be defined as proto messages.
The equivalent Go structs will then be generated, along with utilities to
validate, publish, and consume messages.

## Defining a simple event

If not told otherwise, `protoc-gen-event` considers every proto message ending with `Event` to be an event.

```protobuf
// example.proto
syntax = "proto3";

package examplepb.v1;

option go_package = "github.com/myorg/protoc-gen-event-example/example";

import "google/protobuf/timestamp.proto";

message UserRegisteredEvent {
  string accountID = 1;
  google.protobuf.Timestamp registeredAt = 2;
}
```

The code for the event above can be generated with:

```{.shell hl_lines="2-3"}
protoc -I . -I /usr/include \
    --event_out=. \
    --event_opt=paths=source_relative \
    --go_out=paths=source_relative:. \
    ./example.proto
```

As a result of this command, you will get:

* a `UserRegisteredEvent` Go struct representing the event
* a `UserRegisteredEvent.Publish(ctx context.Context, publisher message.Publisher)` function, to publish the event using Watermill
* a `UserRegisteredEvent.PublishWithUUID(ctx context.Context, publisher message.Publisher, uuid string)` function, to publish the event with a given UUID
* a `UserRegisteredEventHandler(f func(pe *UserRegisteredEvent, m *message.Message) error)` function, that can be used alongside Watermill to decode and consume events

**Note:** the `Publish*()` functions will automatically publish on a topic named after the event (`examplepb.v1.UserRegisteredEvent` in this case)

## Publishing on a custom topic

By default, the topic to which events will be published is derived from the proto package and event name.

This setting can be overridden with the `topic_name` option:

```{.protobuf hl_lines="9 12"}
// example.proto
syntax = "proto3";

package examplepb.v1;

option go_package = "github.com/myorg/protoc-gen-event-example/example";

import "google/protobuf/timestamp.proto";
import "pkg/options/descriptor.proto";

message UserRegisteredEvent {
  option (voi.event.options).topic_name = "user-registered-event";
  
  string accountID = 1;
  google.protobuf.Timestamp registeredAt = 2;
}
```

## Custom event suffix

`protoc-gen-event` considers very proto message ending with `Event` to be an event.
To change this setting, the `suffixMatch` option can be set when running `protoc`:

```{.shell hl_lines="3"}
protoc -I . -I /usr/include \
    --event_out=. \
    --event_opt=suffixMatch=MyEventSuffix,paths=source_relative \
    --go_out=paths=source_relative:. \
    ./example.proto
```