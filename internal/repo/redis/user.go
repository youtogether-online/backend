package redis

import (
	"context"
	"time"
)

// ContainsKeys of redis by key
func (r *RClient) ContainsKeys(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Exists(ctx, keys...).Result()
}

// SetVariable of redis by key, his value and exploration time
func (r *RClient) SetVariable(ctx context.Context, key string, value any, exp time.Duration) error {
	return r.client.SetEx(ctx, key, value, exp).Err()
}
