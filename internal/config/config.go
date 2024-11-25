package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"local"`
	Storage    string `yaml:"storage_path" env-required:"true"`
	HTTPConfig `yaml:"http_server"`
}

type HTTPConfig struct {
	Address     string        `yaml:"address" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeOut time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// MustLoad Функция которая прочитает файл конфига и заполнит нат файл конфига
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("No Config_Path local.yaml")
	}

	//Проверяем существует ли такой файл
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s is not", configPath)
	}

	//Объявлем объект конфига
	var cfg Config

	//Считываем путь
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("Error read config file: ", err)
	}

	return &cfg
}
