# Enforcing validation rules

This plugin has a native and optional integration with [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate).

When working with [`protoc-gen-validate`](https://github.com/envoyproxy/protoc-gen-validate), all messages generated include a `Validate() error` function which returns the first error encountered during validation.

This function is what `protoc-gen-event` looks for when publishing events. When it is available and returns an error, the publishing process is halted.

## Disabling validation

Validation can be disabled per event with a message-level annotation:

```protobuf
message SomeEvent {
  option (voi.event.options).no_message_validation = true;

  string accountID = 1 [(validate.rules).string.min_len = 1];
  map<string, string> data = 2;
}

message OtherEvent {
  string accountID = 1 [(validate.rules).string.min_len = 1];
  map<string, string> data = 2;
}
```

In schema above, the first even will not be automatically validated on publish while the second will.