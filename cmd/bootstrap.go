package cmd

import (
	"github.com/spf13/cobra"

	"crgo/config"
	"crgo/infra/conf"
	"crgo/infra/db"
	"crgo/infra/log"
	"crgo/infra/rabbitmq"
	"crgo/infra/redis"
)

func init() {
	// 加载配置文件
	cobra.OnInitialize(conf.InitConfig)
	// 配置函数...
	cobra.OnInitialize(config.Init)
	cobra.OnInitialize(log.InitLogger)
	cobra.OnInitialize(db.InitDB)
	cobra.OnInitialize(redis.InitRedis)
	cobra.OnInitialize(rabbitmq.Init)
}
