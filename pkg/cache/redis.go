package cache

import "time"

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

func (cache *redisCache) Set(key string, value string) {
	
}

func (cache *redisCache) Get(key string, value string) {
	
}