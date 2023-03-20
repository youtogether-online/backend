package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
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
	for k, v := range info {
		if err := r.client.HSetNX(ctx, sessionId, k, v).Err(); err != nil {
			return err
		}
	}
	logrus.Info(r.client.HGetAll(ctx, sessionId).Val())

	return r.client.Expire(ctx, sessionId, month).Err()
}

// GetSession and all its parameters
func (r *RClient) GetSession(ctx context.Context, sessionId string) (map[string]string, error) {
	return r.client.HGetAll(ctx, sessionId).Result()
}

// ExpandExpireSession if key exists and have lesser than 15 days of expire
func (r *RClient) ExpandExpireSession(ctx context.Context, sessionId string) error {
	if v, err := r.client.TTL(ctx, sessionId).Result(); v <= cfg.Session.Duration/2 {
		return r.client.ExpireLT(ctx, sessionId, month).Err()
	} else {
		return err
	}
}

// FindSessionsByUsername returns all existing sessions by username
func (r *RClient) FindSessionsByUsername(ctx context.Context, userName string) []map[string]string {
	res := make([]map[string]string, 5)
	for _, v := range r.client.Keys(ctx, "/^[0-9A-F]{8}-[0-9A-F]{4}-[4][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i").Val() {
		if r.client.HGet(ctx, v, "username").Val() == userName {
			res = append(res, r.client.HGetAll(ctx, v).Val())
		}
	}
	return res
}

// DelSession fully deletes session id
func (r *RClient) DelSession(ctx context.Context, sessionId string) {
	r.client.Del(ctx, sessionId)
}

// EqualsPopCode returns true if code is involved in email and deletes it
func (r *RClient) EqualsPopCode(ctx context.Context, email string, code string) (bool, error) {
	exist, err := r.client.SIsMember(ctx, email, code).Result()
	if err != nil {
		return false, err
	}
	return exist, r.client.Del(ctx, email).Err()
}

// SetCodes or add it to existing key
func (r *RClient) SetCodes(ctx context.Context, key string, value ...any) error {
	err := r.client.SAdd(ctx, key, value...).Err()
	if err != nil {
		return err
	}
	return r.client.Expire(ctx, key, time.Hour).Err()
}
