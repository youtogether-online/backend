package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
)

// Open redis connection and check it. Returns the client of defined redis database
func Open(host string, port, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: net.JoinHostPort(host, strconv.Itoa(port)),
		DB:   db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		logrus.WithError(err).Fatal("Unable to connect to the redis database")
	}

	return client
}
