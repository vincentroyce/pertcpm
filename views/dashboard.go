package views

import (
	"net/http"
	"strings"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/dashboard/")
	return map[string]interface{}{
		"Title": "User Dashboard",
	}
}
