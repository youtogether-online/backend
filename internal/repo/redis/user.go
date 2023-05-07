package redis

import (
	"context"
	"time"
)

func (r *RClient) ContainsKeys(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Exists(ctx, keys...).Result()
}

func (r *RClient) SetVariable(ctx context.Context, key string, value any, exp time.Duration) error {
	return r.client.SetEx(ctx, key, value, exp).Err()
}
