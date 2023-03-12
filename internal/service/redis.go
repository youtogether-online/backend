package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RClient struct {
	client *redis.Client
}

func NewRClient(client *redis.Client) *RClient {
	return &RClient{client: client}
}

// SetVariable or add it in redis database
func (u RClient) SetVariable(key string, value ...any) error {
	err := u.client.SAdd(context.Background(), key, value...).Err()

	u.client.Expire(context.Background(), key, time.Hour)
	return err
}

// GetVariables of redis with key
func (u RClient) GetVariables(key string) ([]string, error) {
	return u.client.SMembers(context.Background(), key).Result()
}

// ContainsVariable return true if value is involved in key
func (u RClient) ContainsVariable(key string, code string) bool {
	return u.client.SIsMember(context.Background(), key, code).Val()
}
