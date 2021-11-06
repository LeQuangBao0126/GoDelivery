package pubsub

import (
	"context"
	"fmt"
	"time"
)

type Topic string

type Pubsub interface {
	Publish(ctx context.Context, channel Topic, data *Message) error
	Subscribe(ctx context.Context, channel Topic) (ch <-chan *Message, close func())
}

type Message struct {
	id        string
	channel   Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC().UnixNano()
	return &Message{
		id:        fmt.Sprintf("%d", now),
		data:      data,
		createdAt: time.Now().UTC(),
	}
}
func (evt *Message) String() string {
	return fmt.Sprintf("Message %s", evt.channel)
}
func (evt *Message) Channel() Topic {
	return evt.channel
}
func (evt *Message) SetChannel(topic Topic) {
	evt.channel = topic
}

func (evt *Message) Data() interface{} {
	return evt.data
}
