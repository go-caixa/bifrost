package config

import (
	"context"
	"fmt"

	"github.com/go-caixa/bifrost/common/logger"
	"github.com/spf13/viper"
)

type Configuration struct {
	Env  string `mapstructure:"env"`
	Port int    `mapstructure:"port"`
}

func (c Configuration) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

func ReadConfig(ctx context.Context, env string) *Configuration {
	var conf Configuration

	fileName := fmt.Sprintf("config.%s.yaml", env)

	logger.Infof(ctx, "reading configuration file %s", fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../")
	viper.SetConfigFile(fileName)

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf(ctx, err, "viper.ReadInConfig return an error")
		return nil
	}

	if err := viper.Unmarshal(&conf); err != nil {
		logger.Fatalf(ctx, err, "viper.Unmarshal retun an error")
		return nil
	}

	logger.Infof(ctx, "reading configuration file %s succeed", fileName)
	return &conf
}
