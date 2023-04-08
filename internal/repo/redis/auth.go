package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/wtkeqrf0/you_together/pkg/conf"
	"time"
)

const month = time.Hour * 24 * 30

var cfg = conf.GetConfig()

type RClient struct {
	client *redis.Client
}

func NewRClient(client *redis.Client) *RClient {
	return &RClient{client: client}
}

// SetSession and all its parameters
func (r *RClient) SetSession(ctx context.Context, sessionId string, info map[string]string) error {
	return r.client.Watch(ctx, func(tx *redis.Tx) error {
		for k, v := range info {
			if err := tx.HSetNX(ctx, sessionId, k, v).Err(); err != nil {
				return err
			}
		}
		return tx.Expire(ctx, sessionId, month).Err()
	}, sessionId)
}

// GetSession and all its parameters
func (r *RClient) GetSession(ctx context.Context, sessionId string) (map[string]string, error) {
	return r.client.HGetAll(ctx, sessionId).Result()
}

// ExpandExpireSession if key exists and have lesser than 15 days of expire
func (r *RClient) ExpandExpireSession(ctx context.Context, sessionId string) (bool, error) {
	v, err := r.client.TTL(ctx, sessionId).Result()
	if v <= cfg.Session.Duration/2 {
		return r.client.Expire(ctx, sessionId, month).Result()
	}
	return false, err
}

// DelKeys fully deletes session id
func (r *RClient) DelKeys(ctx context.Context, keys ...string) {
	r.client.Del(ctx, keys...)
}

// EqualsPopCode returns true if code is involved in email and deletes it
func (r *RClient) EqualsPopCode(ctx context.Context, email string, code string) (exist bool, err error) {
	err = r.client.Watch(ctx, func(tx *redis.Tx) error {
		exist, err = tx.SIsMember(ctx, email, code).Result()
		if err != nil {
			return err
		}
		return tx.Del(ctx, email).Err()
	}, email)

	return
}

// SetCodes or add it to existing key
func (r *RClient) SetCodes(ctx context.Context, key string, value ...any) error {
	return r.client.Watch(ctx, func(tx *redis.Tx) error {
		err := tx.SAdd(ctx, key, value...).Err()
		if err != nil {
			return err
		}
		return tx.Expire(ctx, key, time.Hour).Err()
	}, key)
}
