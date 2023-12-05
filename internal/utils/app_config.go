package utils

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type appConfig struct {
	Environment       string `mapstructure:"ENVIRONMENT"`
	DbConnection      string `mapstructure:"DB_CONNECTION"`
	HTTPServerAddress string `mapstructure:"HTTPServerAddress"`
	ContextTimeout    int    `mapstructure:"CONTEXT_TIMEOUT"`
}

// LoadAppConfig reads configuration from file or environment variables.
func LoadAppConfig(path string) (appConfig appConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&appConfig)
	return appConfig, err
}
