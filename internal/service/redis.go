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

// SetCodes or add it to user email
func (u RClient) SetCodes(key string, value ...any) error {
	err := u.client.SAdd(context.Background(), key, value...).Err()

	u.client.Expire(context.Background(), key, time.Hour)
	return err
}

// ContainsKeys of redis by key
func (u RClient) ContainsKeys(keys ...string) (int64, error) {
	return u.client.Exists(context.Background(), keys...).Result()
}

// SetVariable of redis by key, his value and exploration time
func (u RClient) SetVariable(key string, value any, exp time.Duration) error {
	return u.client.SetEx(context.Background(), key, value, exp).Err()
}

// ContainsPopCode return true if code is involved in email and deletes it
func (u RClient) ContainsPopCode(email string, code string) bool {
	b := u.client.SIsMember(context.Background(), email, code).Val()
	u.client.Del(context.Background(), email)
	return b
}
