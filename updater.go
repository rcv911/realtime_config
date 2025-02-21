package realtime_config

import (
	"log"

	"github.com/rcv911/realtime_config/config"
)

// updateConfig обновляет конфигурацию безопасным образом
func (rt *RealTimeConfig) updateConfig(newConfig *config.Config) {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	rt.config = newConfig
	log.Printf("Config updated: %v \n", rt.config)
}
