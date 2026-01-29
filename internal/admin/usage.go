package admin

import (
	"encoding/json"
	"net/http"

	"choubis/internal/tenant"
)

func serveUsage(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(400)
		return
	}

	u := tenant.Get(id)

	json.NewEncoder(w).Encode(map[string]int{
		"usage": u,
	})
}
