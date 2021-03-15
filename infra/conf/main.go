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


func fromEnv() {
	// viper.SetEnvPrefix("cr")
	viper.SetDefault("env", os.Getenv("env"))
	viper.AutomaticEnv()
	env = strings.ToUpper(viper.GetString("env"))
	if !(IsDev() || IsFP() || IsBeta() || IsProd() || IsTest()) {
		env = "DEV"
	}
	log.Printf("Environment: %s", env)
}

func fromConfigFile() {
	if IsTest() {
		configName = "test"
	} else if !IsDev() {
		configName = "config"
	}
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	log.Printf("Configuration: %s.toml", configName)
	if err := viper.ReadInConfig(); err != nil {
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