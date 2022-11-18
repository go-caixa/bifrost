package config

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-caixa/bifrost/common/logger"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Configuration struct {
	Env      string   `mapstructure:"env"`
	Port     int      `mapstructure:"port"`
	Database Database `mapstructure:"database"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
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

func SetupDBConnection(ctx context.Context, conf Configuration) *sql.DB {
	datasource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.Password, conf.Database.Name)

	logger.Infof(ctx, "connecting to db %s (%s:%s) ...", conf.Database.Name, conf.Database.Host, conf.Database.Port)
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		logger.Fatalf(ctx, err, "conecting to db %s failed", conf.Database.Name)
		return nil
	}

	if err := db.Ping(); err != nil {
		logger.Fatalf(ctx, err, "ping to db %s failed", conf.Database.Name)
		return nil
	}

	logger.Infof(ctx, "connecting to db %s (%s:%s) ...", conf.Database.Name, conf.Database.Host, conf.Database.Port)
	return db
}
