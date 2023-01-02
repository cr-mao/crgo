package conf

import (
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func IsDev() bool {
	return env == "DEV"
}

func IsFP() bool {
	return env == "FP"
}

func IsBeta() bool {
	return env == "BETA"
}

func IsProd() bool {
	return env == "PROD"
}

func IsTest() bool {
	return env == "TEST"
}

var (
	configName = "config.local"
	env        = "DEV"
	viperObj   *viper.Viper
)

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

var NacCosConfigOption NacosConfig

//type networkOption struct {
//	HTTP_ADDR    string
//	HTTP_PORT    string
//	HTTP_CONNECT string
//	GRPC_ADDR    string
//	GRPC_PORT    string
//	GRPC_CONNECT string
//}
//
//var Net networkOption
//
//func defaultNetworkOption() {
//	viper.SetDefault("http_addr", "127.0.0.1")
//	viper.SetDefault("http_port", "8000")
//	viper.SetDefault("http_connect", "127.0.0.1:8000")
//	viper.SetDefault("grpc_addr", "127.0.0.1")
//	viper.SetDefault("grpc_port", "8081")
//	viper.SetDefault("grpc_connect", "127.0.0.1:8081")
//}
//
//func setNetworkOption() {
//	err := viper.Unmarshal(&Net)
//	if err != nil {
//		panic(err)
//	}
//	log.Printf("network config is %#v\n", Net)
//}

//解析nacos配置
func setNacos() {
	instanceConfig := viperObj.GetStringMap("nacos")
	err := mapstructure.Decode(instanceConfig, &NacCosConfigOption)
	if err != nil {
		panic(err)
	}
}

func fromEnv() {
	// viper.SetEnvPrefix("cr")
	viperObj = viper.New()
	viperObj.SetDefault("env", os.Getenv("env"))
	viperObj.AutomaticEnv()
	env = strings.ToUpper(viperObj.GetString("env"))
	if !(IsDev() || IsFP() || IsBeta() || IsProd() || IsTest()) {
		env = "DEV"
	}
	log.Printf("Environment: %s", env)
}

func fromConfigFile() {
	if !IsDev() {
		configName = "config.local"
	}
	viperObj.SetConfigType("toml")
	viperObj.SetConfigName(configName)
	viperObj.AddConfigPath(".")
	log.Printf("Configuration: %s.toml", configName)
	if err := viperObj.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Configuration Not Found")
		} else {
			// config file was found but another error was produced
			panic(err)
		}
	}

	// 监控 .env 文件，变更时重新加载
	viperObj.WatchConfig()
}

func init() {
	OnInitialize(fromEnv, fromConfigFile)
	AfterInit(setNacos)
}

func InitConfig() {
	for _, i := range initializers {
		i()
	}
	for _, a := range after {
		a()
	}

	//fmt.Println(viper.GetStringMap("database"))
}

//获得主 viper对象
func GetViper() *viper.Viper {
	return viperObj
}
