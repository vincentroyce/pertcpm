package api

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func Main(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")

	switch {
	case strings.HasPrefix(r.URL.Path, "/add-project"):
		AddProject(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "Invalid API endpoint",
		})
	}
}
