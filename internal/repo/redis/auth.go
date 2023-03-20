package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const month = time.Hour * 24 * 30

type RClient struct {
	client *redis.Client
}

func NewRClient(client *redis.Client) *RClient {
	return &RClient{client: client}
}

// SetSession and all its parameters
func (r *RClient) SetSession(ctx context.Context, sessionId string, info map[string]string) error {
	err := r.client.HMSet(ctx, sessionId, "username", info["username"], "ip", info["ip"]).Err()
	r.client.Expire(ctx, sessionId, month)
	return err
}

// GetSession and all its parameters
func (r *RClient) GetSession(ctx context.Context, sessionId string) (map[string]string, error) {
	return r.client.HGetAll(ctx, sessionId).Result()
}

// DelSession fully deletes session_id
func (r *RClient) DelSession(ctx context.Context, sessionId string) error {
	return r.client.Del(ctx, sessionId).Err()
}

// EqualsPopCode returns true if code is involved in email and deletes it
func (r *RClient) EqualsPopCode(ctx context.Context, email string, code string) bool {
	b := r.client.SIsMember(ctx, email, code).Val()
	r.client.Del(ctx, email)
	return b
}

// SetCodes or add it to key
func (r *RClient) SetCodes(ctx context.Context, key string, value ...any) error {
	err := r.client.SAdd(ctx, key, value...).Err()
	r.client.Expire(ctx, key, time.Hour)
	return err
}

// GetTTL returns the remaining lifetime of session
func (r *RClient) GetTTL(ctx context.Context, sessionId string) time.Duration {
	return r.client.TTL(ctx, sessionId).Val()
}
