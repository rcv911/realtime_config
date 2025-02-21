package realtime_config

import (
	"context"
	"log"

	"github.com/rcv911/realtime_config/config"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
)

// watchConfigChanges следит за изменениями конфигурации в etcd
func (rt *RealTimeConfig) watchConfigChanges() {
	ctx, cancel := context.WithCancel(context.Background())
	rt.watchCancel = cancel

	watchChan := rt.client.Watch(ctx, rt.configKey)

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			if event.Type == etcdv3.EventTypePut {

				// todo: новый конфиг
				newConfig := &config.Config{}

				if err := yaml.Unmarshal(event.Kv.Value, newConfig); err != nil {
					log.Printf("failed to unmarshal new config: %v \n", err)
					continue
				}

				rt.updateConfig(newConfig)
				log.Printf("updated config: %v \n", newConfig)
			}
		}
	}
}
