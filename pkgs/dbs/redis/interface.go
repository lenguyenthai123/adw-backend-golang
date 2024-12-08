package redis

import (
	"context"
	goredis "github.com/go-redis/redis/v8"
	"time"
)

type Redis interface {
	IsConnected() bool
	Get(key string, value interface{}) error
	Set(key string, value interface{}) error
	SetWithExpiration(key string, value interface{}, expiration time.Duration) error
	Remove(keys ...string) error
	Keys(pattern string) ([]string, error)
	RemovePattern(pattern string) error
	Publish(channel string, message interface{}) error
	Subscribe(ctx context.Context, channel string, handler func(msg *goredis.Message)) error
}
