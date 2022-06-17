package conf

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)



func defaultDconfigOption() {
	// NOTE: see https://github.com/spf13/viper/issues/747
}


var (
	configName = "config.local"
)


var viperObj *viper.Viper
func fromEnv() {
	// viper.SetEnvPrefix("cr")
	viperObj=viper.New()
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
}

func init() {
	OnInitialize(fromEnv,fromConfigFile)
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