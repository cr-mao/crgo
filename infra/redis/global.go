package redis

import (
	"crgo/infra/conf"
	"strings"
	"sync"

	goRedis "github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"

	"crgo/infra/log"
)

const (
	REDIS = "redis"
)

type RedisMap struct {
	mapping map[string]*goRedis.Client
}

var redisMap RedisMap
var once sync.Once

func InitRedis() {
	once.Do(func() {
		mapping := make(map[string]*goRedis.Client)
		vip := conf.GetViper()
		for bind := range vip.GetStringMap(REDIS) {
			instanceConfig := vip.GetStringMap(strings.Join([]string{REDIS, bind}, "."))

			var option goRedis.Options
			err := mapstructure.Decode(instanceConfig, &option)
			if err != nil {
				panic(err)
			}
			log.Debugf("preparing Redis redis.Client -> %s @ %s:%d", bind, option.Addr, option.DB)
			mapping[bind] = goRedis.NewClient(&option)
		}

		redisMap = RedisMap{mapping: mapping}
	})
}
