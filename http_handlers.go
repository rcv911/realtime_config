package realtime_config

import (
	"encoding/json"
	"net/http"

	"github.com/rcv911/realtime_config/config"
	"gopkg.in/yaml.v3"
)

func (rt *RealTimeConfig) GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	rt.mutex.RLock()
	defer rt.mutex.RUnlock()

	val, err := rt.client.Get(r.Context(), rt.configKey)
	if err != nil {
		rt.logger.Error().Stack().Err(err).Send()
		return
	}

	cfg := &config.Config{}
	err = yaml.Unmarshal(val, cfg)
	if err != nil {
		rt.logger.Error().Stack().Err(err).Send()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(cfg)
	if err != nil {
		rt.logger.Error().Stack().Err(err).Send()
		return
	}

	rt.logger.Info().Msgf("%d | get config", http.StatusOK)
}

func (rt *RealTimeConfig) UpdateConfigHandler(w http.ResponseWriter, r *http.Request) {
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		rt.logger.Error().Msgf("%d | invalid request body", http.StatusBadRequest)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	if newValue, ok := updates["TmpStr"].(string); ok {
		rt.config.TmpStr = newValue
	}
	if newValue, ok := updates["TmpInt"].(float64); ok { // JSON numbers decode as float64
		rt.config.TmpInt = int(newValue)
	}

	updatedConfig, err := yaml.Marshal(rt.config)
	if err != nil {
		rt.logger.Error().Msgf("%d | failed to marshal config", http.StatusInternalServerError)
		http.Error(w, "failed to marshal config", http.StatusInternalServerError)
		return
	}

	if err = rt.client.Put(r.Context(), rt.configKey, string(updatedConfig)); err != nil {
		rt.logger.Error().Msgf("%d | failed to update config in etcd", http.StatusInternalServerError)
		http.Error(w, "failed to update config in etcd", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	rt.logger.Info().Msgf("%d | update config", http.StatusOK)
}
