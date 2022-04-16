package cmd

import (
	"github.com/spf13/cobra"

	"crgo/infra/conf"
	"crgo/infra/db"
	"crgo/infra/log"
	"crgo/infra/redis"
)

func init() {
	cobra.OnInitialize(conf.InitConfig)
	cobra.OnInitialize(log.InitLogger)
	cobra.OnInitialize(db.InitDB)
	cobra.OnInitialize(redis.InitRedis)
	//cobra.OnInitialize(rabbitmq.Init)
}
