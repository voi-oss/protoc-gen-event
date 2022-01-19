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

	handler := examples.NotifyEventHandler(func(pe *examples.NotifyEvent, m *message.Message) error {
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

	handler.Handle(<-messages)
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

	handler.Handle(<-messages)
}
