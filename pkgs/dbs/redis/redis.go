package redis

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"backend-golang/pkgs/log"
	goredis "github.com/go-redis/redis/v8"
)

// Timeout is the default timeout for redis operations
const (
	Timeout = 1
)

// Instance is a singleton instance of the Redis struct
var Instance *redis
var once sync.Once

// redis is a struct that implements the Redis interface.
type redis struct {
	client *goredis.Client
}

// NewRedis creates a new Redis instance and returns a Redis interface.
// This function should only be called once.
func NewRedis(connection Connection) Redis {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
		defer cancel()

		rdb := goredis.NewClient(&goredis.Options{
			Addr:     connection.Address,
			Password: connection.Password,
			DB:       connection.Database,
		})

		pong, err := rdb.Ping(ctx).Result()
		if err != nil {
			log.JsonLogger.Error(pong, err)
			Instance = nil
			return
		}

		Instance = &redis{
			client: rdb,
		}
	})

	return Instance
}

// IsConnected checks if the redis connection is alive.
func (r *redis) IsConnected() bool {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	if r.client == nil {
		return false
	}

	_, err := r.client.Ping(ctx).Result()

	return err == nil
}

// Get gets the value of a key and stores it in the value pointer.
func (r *redis) Get(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	strValue, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(strValue), value)
	if err != nil {
		return err
	}

	return nil
}

// SetWithExpiration sets the value of a key with an expiration time.
func (r *redis) SetWithExpiration(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	bData, _ := json.Marshal(value)
	err := r.client.Set(ctx, key, bData, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// Set sets the value of a key.
func (r *redis) Set(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	bData, _ := json.Marshal(value)
	err := r.client.Set(ctx, key, bData, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// Remove removes a key from redis.
func (r *redis) Remove(keys ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	err := r.client.Del(ctx, keys...).Err()
	if err != nil {
		return err
	}

	return nil
}

// Keys returns all keys matching a pattern.
func (r *redis) Keys(pattern string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

// RemovePattern removes all keys matching a pattern.
func (r *redis) RemovePattern(pattern string) error {
	keys, err := r.Keys(pattern)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	err = r.Remove(keys...)
	if err != nil {
		return err
	}

	return nil
}

// Publish publishes a message to a channel.
func (r *redis) Publish(channel string, message interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	bData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return r.client.Publish(ctx, channel, bData).Err()
}

// Subscribe subscribes to a channel and invokes the provided handler for each message received.
func (r *redis) Subscribe(ctx context.Context, channel string, handler func(msg *goredis.Message)) error {
	sub := r.client.Subscribe(ctx, channel)

	defer func(sub *goredis.PubSub) {
		err := sub.Close()
		if err != nil {
			return
		}
	}(sub)

	ch := sub.Channel()

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-ch:
			if !ok {
				return nil
			}
			handler(msg)
		}
	}
}
