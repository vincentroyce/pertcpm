package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/index/")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	uadmin.RenderHTML(w, r, "templates/index.html", nil)
}
