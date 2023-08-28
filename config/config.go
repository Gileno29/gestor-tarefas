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
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config")
	//viper.AddConfigPath("/home/gileno/documents/gestor-tarefas/config/")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	cfg = new(Config)
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	fmt.Println("Todas as chaves do viper", viper.AllKeys(), viper.AllSettings())
	fmt.Println("Essa é minha porta: ", viper.GetString("database"))

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	fmt.Println("Configurações do dabasee: ", cfg.DB)

	return nil

}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {

	return cfg.API.Port

}
