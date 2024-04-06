package environment

import (
	"fmt"

	validator "github.com/asaskevich/govalidator"
	"github.com/spf13/viper"
)

type EnvLoaderParams struct {
	File string
	Path string
	Type string
}

type envLoader struct{}

func NewEnvLoader(params *EnvLoaderParams) error {
	if params == nil {
		params = &EnvLoaderParams{
			File: ".env",
			Path: ".",
			Type: "env",
		}
	}

	envLoader := envLoader{}

	_, err := envLoader.Load(params)
	if err != nil {
		return err
	}

	return nil
}

func (e *envLoader) Load(params *EnvLoaderParams) (*EnvSchema, error) {
	viper.SetConfigName(fmt.Sprintf("%s_config", params.Type))
	viper.SetConfigType(params.Type)
	viper.SetConfigFile(params.File)
	viper.AddConfigPath(params.Path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&envSchema)
	if err != nil {
		return nil, err
	}

	if err := e.validate(); err != nil {
		return nil, err
	}

	return &envSchema, nil
}

func (e *envLoader) validate() error {
	_, err := validator.ValidateStruct(&envSchema)
	if err != nil {
		return err
	}

	return nil
}
