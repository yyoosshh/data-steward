package config

import "github.com/spf13/viper"

type Config struct {
	Port     string
	DBDriver string
	DBSource string
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.SetDefault("port", "8080")
	viper.SetDefault("db_driver", "mysql")
	viper.SetDefault("db_source", "root:password@tcp(localhost:3306)/data_steward?parseTime=true")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
