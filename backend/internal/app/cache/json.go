package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// GetJSON retrieves a JSON value from Redis and unmarshals it into dest.
func GetJSON(ctx context.Context, client *redis.Client, key string, dest interface{}) error {
	data, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// SetJSON marshals value to JSON and stores it in Redis with the given TTL.
func SetJSON(ctx context.Context, client *redis.Client, key string, value interface{}, ttl time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, b, ttl).Err()
}
