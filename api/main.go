package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func Main(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")

	switch {
	case strings.HasPrefix(r.URL.Path, "/add-project"):
		AddProject(w, r)
	case strings.HasPrefix(r.URL.Path, "/delete-project"):
		DeleteProjectAPI(w, r)
	case strings.HasPrefix(r.URL.Path, "/set-predecessor"):
		SetPredecessorAPI(w, r)
	case strings.HasPrefix(r.URL.Path, "/complete-project"):
		CompleteProjectAPI(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "Invalid API endpoint",
		})
	}
}

func CompleteProjectAPI(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	project := models.Project{}
	err := uadmin.Get(&project, "id = ?", id)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"err_msg": "unable to get the project. " + err.Error(),
			"status":  "error",
		})
		return
	}

	project.Completed = true
	now := time.Now()
	project.CompletedAt = &now
	project.Save()

	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"response": "Project set to completed.",
		"status":   "ok",
	})
}
