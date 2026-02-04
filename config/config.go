package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the application configuration.
type Config struct {
	LogLevel     string `mapstructure:"log_level"`
	LedgerDBPath string `mapstructure:"ledger_db_path"`
	APIPort      string `mapstructure:"api_port"`
}

// LoadConfig reads configuration from file and environment variables.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &cfg, nil
}
