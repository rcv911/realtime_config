package realtime_config

import (
	"context"
	"sync"

	"github.com/rcv911/realtime_config/config"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

type ETCDClient interface {
	Get(ctx context.Context, key string) error
	Put(ctx context.Context, key, val string) error
	Watch(ctx context.Context, key string) etcdv3.WatchChan
	Close() error
}

// RealTimeConfig представляет структуру для работы с real-time конфигурацией
type RealTimeConfig struct {
	client      ETCDClient
	config      *config.Config
	mutex       sync.RWMutex
	configKey   string
	watchCancel context.CancelFunc
}

// NewRealTimeConfig инициализирует соединение с etcd и загружает начальную конфигурацию
func NewRealTimeConfig(etcdClient ETCDClient, configKey string) (*RealTimeConfig, error) {
	return &RealTimeConfig{
		client:    etcdClient,
		config:    &config.Config{}, // todo: явная зависимость. map?
		configKey: configKey,
	}, nil
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
