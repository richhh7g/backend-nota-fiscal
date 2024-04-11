package environment

import (
	"errors"

	"github.com/spf13/viper"
)

var (
	ErrNotFoundKey = errors.New("key not found")
)

func Get(key string) (interface{}, error) {
	if !viper.IsSet(key) {
		return nil, ErrNotFoundKey
	}

	return viper.Get(key), nil
}

func GetAll() *EnvSchema {
	return &envSchema
}
