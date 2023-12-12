package utils

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type appConfig struct {
	Environment          string `mapstructure:"ENVIRONMENT"`
	DbConnection         string `mapstructure:"DB_CONNECTION"`
	HTTPServerAddress    string `mapstructure:"HTTPServerAddress"`
	ContextTimeout       int    `mapstructure:"CONTEXT_TIMEOUT"`
	AccessTokenDuration  int    `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration int    `mapstructure:"REFRESH_TOKEN_DURATION"`
	AccessTokenSecret    string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret   string `mapstructure:"REFRESH_TOKEN_SECRET"`
	EmailSenderName      string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

var AppConfig *appConfig

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
	err = viper.Unmarshal(&AppConfig)
	err = viper.Unmarshal(&appConfig)
	return appConfig, err
}
