package redis

import (
	"context"
	redisClient "github.com/go-redis/redis/v8"
)

type subscription struct {
	ctx        context.Context
	channel    string
	subscriber *redisClient.PubSub
}

func newSubscription(id, channel string, client *redisClient.Client) *subscription {
	ctx := context.WithValue(context.Background(), "id", id)
	subscriber := client.Subscribe(ctx, channel)

	return &subscription{
		ctx:        ctx,
		channel:    channel,
		subscriber: subscriber,
	}
}

func (s *subscription) close() error {
	return s.subscriber.Unsubscribe(s.ctx, s.channel)
}
