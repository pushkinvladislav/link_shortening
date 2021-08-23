package config

import (
	"github.com/pushkinvladislav/link_shortening/utils"
	"github.com/spf13/viper"
)

const Directory = "./config"

func getConfigFiles() []string {
	return []string{"config"}
}

func Init() {
	viper.AddConfigPath(Directory)

	for _, filePath := range getConfigFiles() {
		viper.SetConfigName(filePath)
		err := viper.MergeInConfig()
		if err != nil {
			logger.Logger.Fatalf("failed to open config file: %s", err.Error())
		}
	}
}
