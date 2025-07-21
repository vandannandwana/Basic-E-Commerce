package config

import (
	"flag"
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configration file")
		flag.Parse()
		configPath = *flags

		if configPath == ""{
			log.Fatal("Config Path not found")
		}

	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil{
		log.Fatalf("Cannot find the config file: %s", err.Error())
	}

	return &cfg

}
