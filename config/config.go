package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config представляет конфигурацию приложения
type Config struct {
	TmpStr string `yaml:"TMP_STR"`
	TmpInt int    `yaml:"TMP_INT"`
}

// todo: atomic ?

// LoadConfigFromFile загружает конфигурацию из YAML файла (например, на случай initial-запуска)
func LoadConfigFromFile(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := &Config{}
	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// todo: в мапу для универсальности?
	return config, nil
}
