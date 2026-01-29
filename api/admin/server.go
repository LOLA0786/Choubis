package admin

import (
	"encoding/json"
	"net/http"

	"choubis/internal/policy"
)

func Handler() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/policies", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			policy.Load("policies/rules.yaml")

			w.Write([]byte(`{"status":"reloaded"}`))
		}

		if r.Method == "GET" {
			json.NewEncoder(w).Encode("ok")
		}
	})

	return mux
}
