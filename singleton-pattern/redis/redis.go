package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var redisInstance *redis.Client
var once sync.Once

func createRedisInstance() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	redisInstance = rdb
}

// GetRedisInstance returns the redis instance
func GetRedisInstance() *redis.Client {
	if redisInstance != nil {
		return redisInstance
	}
	once.Do(createRedisInstance)

	return redisInstance
}
