package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rcv911/realtime_config"
	"github.com/rcv911/realtime_config/pkg/etcd"
	"github.com/rs/zerolog"
	etcdv3 "go.etcd.io/etcd/client/v3"
)

const timeout = 3 * time.Second

func main() {
	ctx := context.Background()

	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	client, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: timeout,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}

	etcdClient := etcd.New(client)

	configKey := "/app/config"
	filepath := "config/config_template.yaml"

	// todo: realtime_config как клиент/либа, которая ходит в etcd + по эвенту (канал?)
	rtConfig, err := realtime_config.NewRealTimeConfig(logger, etcdClient, configKey)
	if err != nil {
		log.Fatalf("failed to initialize real-time config: %v", err)
	}
	defer rtConfig.Close()

	// загрузить тестовый конфиг
	err = rtConfig.LoadConfig(ctx, filepath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	go rtConfig.WatchConfigChanges()

	// поднимаем http сервер
	realtime_config.StartServer(rtConfig)
}
