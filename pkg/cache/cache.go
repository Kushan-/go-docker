package cache

import (
	"time"
	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host string
	cacheDb   int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration)  {
	return &redisCache {
		host : host
		cacheDb: db
		expires: exp
	}	
}

func (cache *redisCache) getClient() *redis.Client  {
	return redis.NewClient(&redis.option) {
		Addr : cache.host
		Password: ""
		DB: 0
	}	
}



func (cache *redisCache) Set(key string, value string) {
	Client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string, value string) {
	Client := cache.getClient()

	json, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	err = json.Unmarshal([]byte(json))
	if err != nil{
		panic(err)
	}

	return json
}