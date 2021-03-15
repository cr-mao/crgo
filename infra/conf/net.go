package conf

import (
	"github.com/spf13/viper"
	"log"
)

type networkOption struct {
	HTTP_ADDR    string
	HTTP_PORT    string
	HTTP_CONNECT string
	GRPC_ADDR    string
	GRPC_PORT    string
	GRPC_CONNECT string
}

var Net networkOption

func defaultNetworkOption() {
	viper.SetDefault("http_addr", "127.0.0.1")
	viper.SetDefault("http_port", "8000")
	viper.SetDefault("http_connect", "127.0.0.1:8000")
	viper.SetDefault("grpc_addr", "127.0.0.1")
	viper.SetDefault("grpc_port", "8081")
	viper.SetDefault("grpc_connect", "127.0.0.1:8081")
}

func setNetworkOption() {
	err := viper.Unmarshal(&Net)
	if err != nil {
		panic(err)
	}
	log.Printf("network config is %#v\n", Net)
}
func init() {
	OnInitialize(defaultNetworkOption)
	AfterInit(setNetworkOption)
}
