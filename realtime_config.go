package realtime_config

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/rcv911/realtime_config/config"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
)

const dialTimeout = 3 * time.Second

// RealTimeConfig представляет структуру для работы с real-time конфигурацией
type RealTimeConfig struct {
	client      *etcdv3.Client
	config      *config.Config
	mutex       sync.RWMutex
	configKey   string
	watchCancel context.CancelFunc
}

// NewRealTimeConfig инициализирует соединение с etcd и загружает начальную конфигурацию
func NewRealTimeConfig(etcdEndpoints []string, configKey string) (*RealTimeConfig, error) {
	client, err := etcdv3.New(etcdv3.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to etcd: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	//resp, err := client.Get(ctx, configKey)
	//cancel()
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get config from etcd: %w", err)
	//}

	// todo: конфиг
	cfg, err := config.LoadConfigFromFile("config/config_template.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to load config from file: %w", err)
	}

	cfgBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal config: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	resp, err := client.Put(ctx, configKey, string(cfgBytes))
	cancel()
	if err != nil {
		log.Fatalf("Failed to put key: %v", err)
	}

	fmt.Println(resp)

	//if err = yaml.Unmarshal(resp.Kvs[0].Value, cfg); err != nil {
	//	return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	//}

	rtConfig := &RealTimeConfig{
		client:    client,
		config:    cfg,
		configKey: configKey,
	}

	go rtConfig.watchConfigChanges()

	return rtConfig, nil
}

// GetConfig безопасно возвращает текущую конфигурацию
func (rt *RealTimeConfig) GetConfig() *config.Config {
	rt.mutex.RLock()
	defer rt.mutex.RUnlock()

	return rt.config
}

// Close завершает соединение с etcd
func (rt *RealTimeConfig) Close() {
	rt.watchCancel()
	rt.client.Close()
}
