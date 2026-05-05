package examples_test

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/stretchr/testify/require"
	"github.com/voi-oss/protoc-gen-event/examples"
)

func TestSimple(t *testing.T) {
	r := require.New(t)

	ne := examples.NotifyEvent{
		AccountID: "some-cool-UUID-1234567890",
		Data: map[string]string{
			"key": "value",
		},
	}

	logger := watermill.NewStdLogger(false, false)

	publisher := gochannel.NewGoChannel(gochannel.Config{
		OutputChannelBuffer: 10,
	}, logger)

	handler := examples.NotifyEventHandler(func(pe *examples.NotifyEvent, _ *message.Message) error {
		r.Equal(pe.AccountID, ne.AccountID)
		return nil
	})
	ctx := context.Background()
	messages, err := publisher.Subscribe(ctx, handler.Topic())
	r.NoError(err)

	if err := ne.Publish(ctx, publisher); err != nil {
		t.Error(err)
		return
	}

	r.NoError(handler.Handle(<-messages))
}

// TestDefaultMarshal_ZeroValueFieldsOmitted verifies that with the default generator
// settings (emitUnpopulated=false), proto3 zero-value fields are absent from the
// published JSON payload. When the generator is invoked with emitUnpopulated=true
// those fields are always present in the output.
func TestDefaultMarshal_ZeroValueFieldsOmitted(t *testing.T) {
	r := require.New(t)

	ne := examples.NotifyEvent{} // AccountID is zero value ""

	logger := watermill.NewStdLogger(false, false)
	publisher := gochannel.NewGoChannel(gochannel.Config{OutputChannelBuffer: 10}, logger)
	ctx := context.Background()

	topic := examples.NotifyEventHandler(nil).Topic()
	messages, err := publisher.Subscribe(ctx, topic)
	r.NoError(err)

	r.NoError(ne.Publish(ctx, publisher))

	msg := <-messages
	r.NotContains(string(msg.Payload), `"accountID"`, "zero-value fields must be absent from JSON payload with default generator settings")
}

func TestSimpleAttribute(t *testing.T) {
	r := require.New(t)

	ne := examples.AttributeEvent{
		AccountID: "some-cool-UUID-1234567890",
		ZoneID:    "123",
		Data: map[string]string{
			"key": "value",
		},
	}

	logger := watermill.NewStdLogger(false, false)

	publisher := gochannel.NewGoChannel(gochannel.Config{
		OutputChannelBuffer: 10,
	}, logger)

	handler := examples.AttributeEventHandler(func(pe *examples.AttributeEvent, m *message.Message) error {
		r.Equal(pe.AccountID, ne.AccountID)
		r.Equal(m.Metadata.Get("AccountID"), "some-cool-UUID-1234567890")
		r.Equal(m.Metadata.Get("zone_id"), "123")
		return nil
	})
	ctx := context.Background()
	messages, err := publisher.Subscribe(ctx, handler.Topic())
	r.NoError(err)

	if err := ne.Publish(ctx, publisher); err != nil {
		t.Error(err)
		return
	}

	r.NoError(handler.Handle(<-messages))
}
