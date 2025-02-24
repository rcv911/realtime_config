package realtime_config

import (
	"log"
	"net/http"
)

func StartServer(rt *RealTimeConfig) {
	if rt == nil {
		log.Fatal("rt config is nil")
	}

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			rt.GetConfigHandler(w, r)
		case http.MethodPut:
			rt.UpdateConfigHandler(w, r)
		default:
			rt.logger.Error().Msgf("method '%s' not allowed", r.Method)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	rt.logger.Info().Msg("starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
