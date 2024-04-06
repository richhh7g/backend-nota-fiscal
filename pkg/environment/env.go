package environment

import (
	"fmt"

	"github.com/spf13/viper"
)

func Get(key string) (interface{}, error) {
	if !viper.IsSet(key) {
		return nil, fmt.Errorf("key %s not found", key)
	}

	return viper.Get(key), nil
}

func GetAll() *EnvSchema {
	return &envSchema
}
