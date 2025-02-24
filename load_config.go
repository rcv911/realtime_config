package realtime_config

import (
	"context"
	"fmt"

	"github.com/rcv911/realtime_config/config"
	"gopkg.in/yaml.v3"
)

// LoadConfig начальная загрузка тестового конфига
func (rt *RealTimeConfig) LoadConfig(ctx context.Context, filepath string) error {
	cfg, err := config.LoadConfigFromFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to load config from file: %w", err)
	}

	cfgBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = rt.client.Put(ctx, rt.configKey, string(cfgBytes))
	if err != nil {
		return fmt.Errorf("failed to put config: %w", err)
	}

	return nil
}
