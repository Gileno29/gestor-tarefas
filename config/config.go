package config

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Dtabase string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database,host", "localhost")
	viper.SetDefault("database.port", "5332")
}

func load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.SetConfigFile(".")
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		return err
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:    viper.GetString("database.host"),
		Port:    viper.GetString("database.port"),
		User:    viper.GetString("database.user"),
		Pass:    viper.GetString("database.pass"),
		Dtabase: viper.GetString("database.name"),
	}

	return nil

}
