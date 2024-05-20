package api

import (
	"fmt"
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func DeleteProjectAPI(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fmt.Println("id: ", id)
	project := models.Project{}
	err := uadmin.Get(&project, "id = ?", id)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":    "error",
			"error_msg": "unable to get the project. " + err.Error(),
		})
		return
	}
	err = uadmin.Delete(&project)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":    "error",
			"error_msg": "unable to delete the project. " + err.Error(),
		})
		return
	}

	phases := []models.Phase{}
	err = uadmin.DeleteList(&phases, "project_id = ?", id)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":    "error",
			"error_msg": "unable to delete the phases from this project. " + err.Error(),
		})
		return
	}

	activity := []models.Activity{}
	err = uadmin.DeleteList(&activity, "project_id = ?", id)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":    "error",
			"error_msg": "unable to delete the activities from this project. " + err.Error(),
		})
		return
	}

	subActivity := []models.SubActivity{}
	err = uadmin.DeleteList(&subActivity, "project_id = ?", id)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":    "error",
			"error_msg": "unable to delete the sub activities from this project. " + err.Error(),
		})
		return
	}

	worker := []models.Worker{}
	err = uadmin.DeleteList(&worker, "project_id = ?", id)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":    "error",
			"error_msg": "unable to delete the workers from this project. " + err.Error(),
		})
		return
	}

	equipment := []models.Equipment{}
	err = uadmin.DeleteList(&equipment, "project_id = ?", id)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":    "error",
			"error_msg": "unable to delete the equipments from this project. " + err.Error(),
		})
		return
	}

	uadmin.ReturnJSON(w, r, map[string]any{
		"status":   "ok",
		"response": "Project Deleted",
	})
}
