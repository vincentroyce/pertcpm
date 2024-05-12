package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func AddProject(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "Invalid HTTP method",
		})
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusPreconditionFailed)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "invalid MIME type",
		})
		return
	}

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "unable to read body." + err.Error(),
		})
		return
	}

	if len(buf) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "empty body",
		})
		return
	}

	obj := map[string]any{}
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "unable to parse JSON body. " + err.Error(),
		})
		return
	}
	project := models.Project{}
	project.Name = fmt.Sprintf("%v", obj["projectName"])
	cost := fmt.Sprintf("%v", obj["cost"])
	project.Cost, _ = strconv.ParseInt(cost, 10, 0)
	project.Save()

	projectObj := obj["obj"].(map[string]any)
	for k, v := range projectObj {
		phaseObj := v.(map[string]any)
		var phaseID uint
		phaseNum := k
		for k, v := range phaseObj {
			phase := models.Phase{}
			phase.No = phaseNum
			if arr, ok := v.([]interface{}); ok {
				phase.ProjectID = project.ID
				phase.Name = k
				phase.OptimisticTime = int(arr[0].(float64))
				phase.MostLikelyTime = int(arr[1].(float64))
				phase.PessimisticTime = int(arr[2].(float64))
				phase.Save()
				phaseID = phase.ID
				continue
			}
			activityObj := v.(map[string]any)
			var activityID uint
			var name string
			actNum := k
			for k, v := range activityObj {
				if arr, ok := v.([]interface{}); ok {
					activity := models.Activity{}
					activity.No = actNum
					activity.ProjectID = project.ID
					activity.PhaseID = phaseID
					activity.Name = k
					activity.OptimisticTime = int(arr[0].(float64))
					activity.MostLikelyTime = int(arr[1].(float64))
					activity.PessimisticTime = int(arr[2].(float64))
					activity.Save()
					activityID = activity.ID
					name = activity.Name
					continue
				}
				subActivityObj := v.(map[string]any)
				subActNum := k
				for k, v := range subActivityObj {

					if arr, ok := v.([]interface{}); ok {
						subActivity := models.SubActivity{}
						subActivity.No = subActNum
						subActivity.ProjectID = project.ID
						subActivity.PhaseID = phaseID
						subActivity.ActivityID = activityID
						subActivity.Name = k
						subActivity.OptimisticTime = int(arr[0].(float64))
						subActivity.MostLikelyTime = int(arr[1].(float64))
						subActivity.PessimisticTime = int(arr[2].(float64))
						if name == "" {
							subActivity.PhaseDirect = true
						}
						subActivity.Save()
					}
				}
			}
		}
	}

	uadmin.ReturnJSON(w, r, map[string]any{
		"status":   "ok",
		"response": obj,
	})
}
