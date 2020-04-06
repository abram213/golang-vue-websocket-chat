package app

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	//SecretKey for password hashing
	SecretKey []byte

	// APIKey for jwt token generation
	APIKey string
}

func InitConfig() (*Config, error) {
	config := &Config{
		SecretKey: []byte(viper.GetString("SecretKey")),
		APIKey:    viper.GetString("APIKey"),
	}
	if len(config.SecretKey) == 0 {
		return nil, fmt.Errorf("SecretKey must be set!")
	}
	if len(config.APIKey) == 0 {
		return nil, fmt.Errorf("APIKey must be set!")
	}
	return config, nil
}
