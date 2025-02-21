package main

import (
	"fmt"
	"log"

	"github.com/rcv911/realtime_config"
)

func main() {
	//etcdEndpoints := []string{"localhost:2379"}
	etcdEndpoints := []string{"http://127.0.0.1:2379"}
	configKey := "/app/config"

	rtConfig, err := realtime_config.NewRealTimeConfig(etcdEndpoints, configKey)
	if err != nil {
		log.Fatalf("failed to initialize real-time config: %v", err)
	}
	defer rtConfig.Close()

	// Инициализируем приложение с начальной конфигурацией
	initialConfig := rtConfig.GetConfig()
	fmt.Printf("Initial config: %+v \n", initialConfig)

	// Приложение работает и динамически обновляет конфигурацию
	select {}
}
