package cache

import (
	"encoding/json"
	"time"

	"com.github/fabiosebastiano/go-rest-api/entity"
	"github.com/go-redis/redis"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

// NewRedisCache
func NewRedisCache(host string, db int, exp time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *entity.Post) {
	redisClient := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	redisClient.Set(key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *entity.Post {
	redisClient := cache.getClient()

	val, err := redisClient.Get(key).Result()
	if err != nil {
		return nil
	}

	post := entity.Post{}
	err = json.Unmarshal([]byte(val), &post)

	if err != nil {
		panic(err)
	}

	return &post
}
