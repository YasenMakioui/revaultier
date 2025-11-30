package server

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		Database string `mapstructure:"database"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}

	Auth struct {
		SecretKey string `mapstructure:"secretkey"`
	}

	Server struct {
		Port int `mapstructure:"port"`
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("revaultier")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/revaultier")
	viper.AddConfigPath("/etc/revaultier")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
