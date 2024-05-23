package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func SetPredecessorAPI(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		handleError(w, r, "unable to parse form. "+err.Error())
		return
	}

	ids := r.FormValue("ids")
	predecessors := r.FormValue("predecessors")
	critical := r.FormValue("critical")

	var idsArr []interface{}
	err = json.Unmarshal([]byte(ids), &idsArr)
	if err != nil {
		handleError(w, r, "Failed to parse idsArr."+err.Error())
		return
	}

	var predecessorArr []interface{}
	err = json.Unmarshal([]byte(predecessors), &predecessorArr)
	if err != nil {
		handleError(w, r, "Failed to parse predecessorArr."+err.Error())
		return
	}

	var criticalArr []interface{}
	err = json.Unmarshal([]byte(critical), &criticalArr)
	if err != nil {
		handleError(w, r, "Failed to parse criticalArr."+err.Error())
		return
	}
	// Batch database operations
	var wg sync.WaitGroup
	for i := range idsArr {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			subActivity := models.SubActivity{}
			project := models.Project{}

			// Fetch subactivity
			err := uadmin.Get(&subActivity, "id = ?", idsArr[i])
			if err != nil {
				handleError(w, r, "unable to find the sub activity from array index "+strconv.Itoa(i)+". "+err.Error())
				return
			}

			// Fetch project
			err = uadmin.Get(&project, "id = ?", subActivity.ProjectID)
			if err != nil {
				handleError(w, r, "unable to get the project. "+err.Error())
				return
			}

			// Fetch predecessors
			predecessors := []string{}
			for _, predID := range predecessorArr[i].([]interface{}) {
				predecessors = append(predecessors, predID.(string))
			}

			// Update subactivity
			subActivity.PredecessorsList = strings.Join([]string(predecessors), ", ")
			subActivity.Critical = criticalArr[i].(bool)
			subActivity.Save()

			// Update project
			project.ScheduleCompleted = true
			project.Save()
		}(i)
	}
	wg.Wait()

	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"status":   "ok",
		"response": "Predecessors have been set to this project",
	})
}

func handleError(w http.ResponseWriter, r *http.Request, errMsg string) {
	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"status":  "error",
		"err_msg": errMsg,
	})
}
