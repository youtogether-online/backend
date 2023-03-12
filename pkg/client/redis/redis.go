package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"net"
)

// Open redis connection and check it. Returns the client of defined redis database
func Open(host, port string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: net.JoinHostPort(host, port),
		DB:   db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		logrus.Fatalf("Unable to connect to the redis database: %v", err)
	}

	return client
}
