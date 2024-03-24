package views

import (
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/home/")
	return map[string]interface{}{
		"Title": "User Home",
	}
}
