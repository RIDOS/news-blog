package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	StorageType string `yaml:"storage_type" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Addr        string        `yaml:"addr" env-required:"true" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5m"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}
