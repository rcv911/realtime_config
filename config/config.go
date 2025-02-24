package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	CfgTmpStr = "TMP_STR"
	CfgTmpInt = "TMP_INT"
)

// LoadConfigFromFile загружает конфигурацию из YAML файла (например, на случай initial-запуска)
func LoadConfigFromFile(filepath string) (map[string]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config map[string]string
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}
