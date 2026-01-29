package admin

import (
	"encoding/json"
	"net/http"

	"choubis/internal/policy"
)

func Handler() http.Handler {

	mux := http.NewServeMux()

	// Reload policies
	mux.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}

		err := policy.Load("policies/rules.yaml")
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"status": "reloaded",
		})
	})

	// Audit log
	mux.HandleFunc("/audit", func(w http.ResponseWriter, r *http.Request) {
		serveAudit(w)
	})

	// Stats
	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		serveStats(w)
	})

	return mux
}
