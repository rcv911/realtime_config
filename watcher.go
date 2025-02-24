package realtime_config

import (
	"context"

	etcdv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
)

// todo: atomic ?

// WatchConfigChanges следит за изменениями конфигурации в etcd
func (rt *RealTimeConfig) WatchConfigChanges() {
	ctx, cancel := context.WithCancel(context.Background())
	rt.watchCancel = cancel

	watchChan := rt.client.Watch(ctx, rt.configKey)

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			if event.Type == etcdv3.EventTypePut {
				var newConfig map[string]string
				if err := yaml.Unmarshal(event.Kv.Value, &newConfig); err != nil {
					rt.logger.Error().Msgf("failed to unmarshal new config: %v", err)
					continue
				}

				rt.updateConfigSelective(newConfig)
				rt.logger.Info().Msgf("updated config: %v", newConfig)
			}
		}
	}
}
