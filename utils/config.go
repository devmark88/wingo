package utils

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// InitConfig : initial application config
func InitConfig(path, prefix string) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix(prefix)
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
}
