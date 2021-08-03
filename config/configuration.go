package config

import "github.com/spf13/viper"

type DatabaseConfig struct {
	UserName     string
	Password     string
	Host         string
	DatabaseName string
	DatabasePort int
}

type MigrationConfig = DatabaseConfig

type ServerConfig struct {
	Port                     int
	GracefullShutdownTimeout int
}

type AppConfig struct {
	Server    ServerConfig
	Database  DatabaseConfig
	Migration MigrationConfig
}

func LoadConfig() (AppConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	var config AppConfig
	err := viper.ReadInConfig()

	if err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
