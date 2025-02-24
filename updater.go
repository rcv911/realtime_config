package realtime_config

import (
	"github.com/rcv911/realtime_config/config"
)

// updateConfig обновляет весь конфиг
func (rt *RealTimeConfig) updateConfig(newConfig *config.Config) {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	rt.config = newConfig

	rt.logger.Info().Msgf("сonfig updated: %v \n", rt.config)
}

// updateConfigSelective обновляет только изменившиеся поля
func (rt *RealTimeConfig) updateConfigSelective(newConfig *config.Config) {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	if rt.config.TmpStr != newConfig.TmpStr {
		rt.logger.Info().Msgf("TmpStr changed: %s -> %s", rt.config.TmpStr, newConfig.TmpStr)
		rt.config.TmpStr = newConfig.TmpStr
	}
	if rt.config.TmpInt != newConfig.TmpInt {
		rt.logger.Info().Msgf("TmpInt changed: %d -> %d", rt.config.TmpInt, newConfig.TmpInt)
		rt.config.TmpInt = newConfig.TmpInt
	}
}
