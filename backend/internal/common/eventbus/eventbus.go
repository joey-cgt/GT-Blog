package event

import (
	"context"
)

type EventBus interface {
	Publish(ctx context.Context, event Event) error
	Subscribe(eventName string, handler EventHandler) error
	UnSubscribe(eventName string) error
}
