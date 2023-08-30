package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

const (
	VERSION     = "version"
	SERVER_PORT = "server.port"
)

func InitViper() {
	// Set the path to your JSON configuration file
	viper.SetConfigFile("config.json")

	// Set the prefix for environment variables
	viper.SetEnvPrefix("PROD")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Enable reading environment variables as well
	viper.AutomaticEnv()

	// Read and parse the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(fmt.Errorf("failed to read configuration file: %s", err))
	}
}
