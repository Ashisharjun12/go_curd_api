package config

import (
	"flag"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
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
            slog.Info("config path is not set")
        }
    }


	if  _,err := os.Stat(configPath); os.IsNotExist(err){
		slog.Info("config file  does not exist", slog.String("configpath", configPath))
	}



	var cfg Config
	err := cleanenv.ReadConfig(configPath,&cfg)
	if err != nil {
        slog.Info("failed to load config:",slog.String("err",err.Error()))
    }

	return &cfg
}