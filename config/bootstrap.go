package config

import "crgo/infra/conf"

func Init(){
	//加载配置
	conf.LoadConfig()
}