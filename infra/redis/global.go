package redis

import (
	"strings"
	"sync"

	goRedis "github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

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

		for bind := range viper.GetStringMap(REDIS) {
			instanceConfig := viper.GetStringMap(strings.Join([]string{REDIS, bind}, "."))

			var option goRedis.Options
			err := mapstructure.Decode(instanceConfig, &option)
			if err != nil {
				panic(err)
			}
			log.Debugf("preparing Redis redis.Client -> %s @ %s", bind, option.Addr)
			mapping[bind] = goRedis.NewClient(&option)
		}

		redisMap = RedisMap{mapping: mapping}
	})
}
