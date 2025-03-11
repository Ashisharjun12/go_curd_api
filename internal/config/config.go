package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string
}

// env-default:"production" set at time of production
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func Mustload() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
        flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
            log.Fatal("config path is not set")
        }
    }


	if  _,err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file %s does not exist", configPath)
	}



	var cfg Config
	err := cleanenv.ReadConfig(configPath,&cfg)
	if err != nil {
        log.Fatalf("failed to load config: %s", err.Error())
    }

	return &cfg
}