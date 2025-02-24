package realtime_config

import "github.com/rcv911/realtime_config/config"

// updateConfig обновляет весь конфиг
func (rt *RealTimeConfig) updateConfig(newConfig map[string]string) {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	rt.config = newConfig

	rt.logger.Info().Msgf("сonfig updated: %+v", rt.config)
}

// updateConfigSelective обновляет только изменившиеся поля
func (rt *RealTimeConfig) updateConfigSelective(newConfig map[string]string) {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	if rt.config[config.CfgTmpStr] != newConfig[config.CfgTmpStr] {
		rt.logger.Info().Msgf("TmpStr changed: %s -> %s", rt.config[config.CfgTmpStr], newConfig[config.CfgTmpStr])
		rt.config[config.CfgTmpStr] = newConfig[config.CfgTmpStr]
	}
	if rt.config[config.CfgTmpInt] != newConfig[config.CfgTmpInt] {
		rt.logger.Info().Msgf("TmpInt changed: %s -> %s", rt.config[config.CfgTmpInt], newConfig[config.CfgTmpInt])
		rt.config[config.CfgTmpInt] = newConfig[config.CfgTmpInt]
	}
}
