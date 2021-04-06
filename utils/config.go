package utils

import "github.com/spf13/viper"

// Config struct stores all configuration of the application
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDR"`
}

func LoadConfig(envDir, configName string) (config Config, err error) {
	viper.AddConfigPath(envDir)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
