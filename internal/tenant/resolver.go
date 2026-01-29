package tenant

import "net/http"

func Resolve(r *http.Request) string {

	t := r.Header.Get("X-Tenant-ID")

	if t == "" {
		return "default"
	}

	return t
}
