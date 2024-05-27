package api

import (
	"net/http"
	"time"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

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
