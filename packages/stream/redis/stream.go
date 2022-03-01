package redis

import (
	"context"
	redisClient "github.com/go-redis/redis/v8"
	"sync"
)

type redisStream struct {
	client        *redisClient.Client
	subscriptions map[string]*subscription
	mutex         *sync.RWMutex
}

func New(Addr, password string) *redisStream {

	client := redisClient.NewClient(&redisClient.Options{
		Addr:     Addr,
		Password: password,
	})

	_, err := client.Ping(context.TODO()).Result()

	if err != nil {
		panic(err)
	}

	return &redisStream{
		client:        client,
		subscriptions: make(map[string]*subscription),
		mutex:         &sync.RWMutex{},
	}
}
