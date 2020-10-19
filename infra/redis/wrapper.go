package redis

import (
	goRedis "github.com/go-redis/redis/v8"
)

const Nil = goRedis.Nil

func Client(bind string) *goRedis.Client {
	return redisMap.mapping[bind]
}
