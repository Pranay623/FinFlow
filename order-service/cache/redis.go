package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	addr := os.Getenv("REDIS_URL")
	if addr == "" {
		addr = "localhost:6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	fmt.Println("Successfully connected to Redis")
}

// CheckIdempotency checks if a key exists in Redis. Returns true if it exists.
func CheckIdempotency(key string) bool {
	exists, _ := RedisClient.Exists(ctx, key).Result()
	return exists > 0
}

// SetIdempotency sets a key in Redis with an expiration
func SetIdempotency(key string, expiration time.Duration) {
	RedisClient.Set(ctx, key, "processed", expiration)
}
