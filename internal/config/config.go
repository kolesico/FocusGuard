package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	*HTTPServer `yaml:"http_server"`
	Env         string `yaml:"env" env-default:"development"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	LogLevel    string `yaml:"log_level" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost"`
	Port        int           `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func LoadConfig(configPath string) (*Config, error) {
	// Must в начале названия функции означает, что если это функция упадет, то аварийно завершит работу.
	// Получаем путь до конфиг файла из env-переменной CONFIG_PATH

	_, err := os.Stat(configPath)
	if err != nil {
		return nil, fmt.Errorf("Error opening config file: %s", err)
	}

	var cfg Config

	// Читаем конфиг файл и заполняем структуру

	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}

	return &cfg, nil
}
