package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func Init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database,host", "localhost")
	viper.SetDefault("database.port", "5332")
}

func Load() error {
	viper.SetConfigType("toml")
	viper.SetConfigName("config.toml")
	//viper.SetConfigFile("./config.toml")
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		fmt.Printf("entre no if")
		return err
	}

	cfg = new(Config)
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	fmt.Println("Essa é minha porta: ", viper.GetInt("api.port"))

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil

}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {

	return cfg.API.Port

}
