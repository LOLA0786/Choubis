package admin

import (
	"encoding/json"
	"net/http"

	"choubis/internal/tenant"
)

func serveTenants(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		var t tenant.Tenant
		json.NewDecoder(r.Body).Decode(&t)

		if t.ID == "" {
			w.WriteHeader(400)
			return
		}

		tenant.AddTenant(t)

		w.Write([]byte(`{"status":"created"}`))
		return
	}

	if r.Method == "GET" {

		all, _ := tenant.LoadAll()

		json.NewEncoder(w).Encode(all)
	}
}
