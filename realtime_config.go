package realtime_config

import (
	"context"
	"sync"

	"github.com/rs/zerolog"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

type ETCDClient interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Put(ctx context.Context, key, val string) error
	Watch(ctx context.Context, key string) etcdv3.WatchChan
	Close() error
}

// RealTimeConfig представляет структуру для работы с real-time конфигурацией
type RealTimeConfig struct {
	logger      zerolog.Logger
	client      ETCDClient
	config      map[string]string
	mutex       sync.RWMutex
	configKey   string
	watchCancel context.CancelFunc
}

// NewRealTimeConfig инициализирует соединение с etcd и загружает начальную конфигурацию
func NewRealTimeConfig(logger zerolog.Logger, etcdClient ETCDClient, configKey string) (*RealTimeConfig, error) {
	return &RealTimeConfig{
		logger:    logger,
		client:    etcdClient,
		config:    make(map[string]string),
		configKey: configKey,
	}, nil
}

// GetConfig безопасно возвращает текущую конфигурацию
func (rt *RealTimeConfig) GetConfig() map[string]string {
	rt.mutex.RLock()
	defer rt.mutex.RUnlock()

	return rt.config
}

// Close завершает соединение с etcd
func (rt *RealTimeConfig) Close() {
	rt.watchCancel()
	rt.client.Close()
}
