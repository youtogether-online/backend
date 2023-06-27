package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"net"
	"strconv"
)

// Open redis connection and check it
func Open(host, password string, port, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(host, strconv.Itoa(port)),
		Password: password,
		DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.WithErr(err).Fatal("unable to connect to the redis database")
	}

	return client
}
