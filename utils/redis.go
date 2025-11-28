package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

// InitRedis initializes the Redis client
func InitRedis() error {
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisPassword := getEnv("REDIS_PASSWORD", "")
	redisDB := 0 // Default DB

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Test connection
	_, err := redisClient.Ping(ctx).Result()
	if err == nil {
		log.Println("Redis connection established successfully")
	}

	return err
}

// StoreUserSession stores user's session with token expiry time
func StoreUserSession(userID string, expiresAt time.Time) error {
	if redisClient == nil {
		if err := InitRedis(); err != nil {
			return err
		}
	}

	// Calculate TTL (time until token expires)
	ttl := time.Until(expiresAt)
	if ttl <= 0 {
		// Token already expired
		return nil
	}

	// Store user session with expiry timestamp
	key := "session:" + userID
	return redisClient.Set(ctx, key, expiresAt.Unix(), ttl).Err()
}

// GetUserSession retrieves user's session expiry time
func GetUserSession(userID string) (time.Time, error) {
	if redisClient == nil {
		if err := InitRedis(); err != nil {
			return time.Time{}, err
		}
	}

	key := "session:" + userID
	val, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// Session doesn't exist
		return time.Time{}, nil
	}
	if err != nil {
		return time.Time{}, err
	}

	// Parse Unix timestamp
	var timestamp int64
	_, parseErr := time.Parse("", val)
	if parseErr == nil {
		return time.Unix(timestamp, 0), nil
	}

	// If parsing fails, try direct conversion
	timestamp = 0
	return time.Unix(timestamp, 0), nil
}

// InvalidateUserSession removes user's session (logout)
func InvalidateUserSession(userID string) error {
	if redisClient == nil {
		if err := InitRedis(); err != nil {
			return err
		}
	}

	key := "session:" + userID
	return redisClient.Del(ctx, key).Err()
}

// IsSessionValid checks if user session is valid and not expired
func IsSessionValid(userID string, tokenExpiresAt time.Time) bool {
	if redisClient == nil {
		if err := InitRedis(); err != nil {
			return false
		}
	}

	key := "session:" + userID
	val, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		// No session found
		return false
	}
	if err != nil {
		// Error checking
		return false
	}

	// Parse stored expiry timestamp
	var storedTimestamp int64
	if _, parseErr := fmt.Sscanf(val, "%d", &storedTimestamp); parseErr != nil {
		return false
	}

	storedExpiry := time.Unix(storedTimestamp, 0)

	// Check if stored session expiry matches token expiry
	// Allow 1 second tolerance for timing differences
	if storedExpiry.Sub(tokenExpiresAt).Abs() > time.Second {
		return false
	}

	// Check if session is not expired
	return time.Now().Before(storedExpiry)
}

// CloseRedis closes the Redis connection
func CloseRedis() error {
	if redisClient != nil {
		return redisClient.Close()
	}
	return nil
}
