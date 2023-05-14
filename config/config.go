package config

import "github.com/spf13/viper"

type Config struct {
	PortRange struct {
		From string `mapstructure:"from"`
		To   string `mapstructure:"to"`
	} `mapstructure:"portRange"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}

	var cfg *Config
	err = viper.Unmarshal(&cfg)
	return cfg, err
}
