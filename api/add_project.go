package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

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

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "unable to parse multipart. " + err.Error(),
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

	if _, ok := obj["projectName"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "missing 'projectName' key in json body",
		})
		return
	}

	if obj["projectName"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "empty project name",
		})
		return
	}

	if _, ok := obj["cost"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "missing 'cost' key in json body",
		})
		return
	}

	if _, ok := obj["description"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "missing 'description' key in json body",
		})
		return
	}

	if obj["description"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "empty description",
		})
		return
	}

	if _, ok := obj["dateStart"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "missing 'dateStart' key in json body",
		})
		return
	}

	if obj["dateStart"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "empty date start",
		})
		return
	}

	if _, ok := obj["dateEnd"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "missing 'dateEnd' key in json body",
		})
		return
	}

	if obj["dateEnd"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "empty date end",
		})
		return
	}

	if _, ok := obj["obj"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "missing 'obj' key in json body",
		})
		return
	}

	if len(obj["obj"].(map[string]any)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "empty project object",
		})
		return
	}

	if _, ok := obj["projectName"].(string); !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "'projectName' json key is not a string",
		})
		return
	}

	if _, ok := obj["cost"].(float64); !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "'cost' json key is not a float",
		})
		return
	}

	if _, ok := obj["description"].(string); !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "'description' json key is not a string",
		})
		return
	}

	if _, ok := obj["dateStart"].(string); !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "'dateStart' json key is not a string",
		})
		return
	}

	if _, ok := obj["dateEnd"].(string); !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "'dateEnd' json key is not a string",
		})
		return
	}

	if _, ok := obj["obj"].(map[string]interface{}); !ok {
		w.WriteHeader(http.StatusBadRequest)
		uadmin.ReturnJSON(w, r, map[string]any{
			"status":  "error",
			"err_msg": "'obj' json key is not a map",
		})
		return
	}

	project := models.Project{}
	project.Name = fmt.Sprintf("%v", obj["projectName"])
	project.Description = fmt.Sprintf("%v", obj["description"])
	cost := fmt.Sprintf("%v", obj["cost"])
	project.Cost, _ = strconv.ParseInt(cost, 10, 0)
	project.DateStart, err = time.Parse(time.RFC3339, fmt.Sprintf("%v", obj["dateStart"]))
	if err != nil {
		fmt.Println("Error parsing start time:", err)
		return
	}
	project.DateEnd, err = time.Parse(time.RFC3339, fmt.Sprintf("%v", obj["dateEnd"]))
	if err != nil {
		fmt.Println("Error parsing end time:", err)
		return
	}
	project.Save()

	projectObj := obj["obj"].(map[string]any)

	// Step 1: Extract and save all phases first
	phaseIDMap := make(map[string]uint)
	for k, v := range projectObj {
		phaseObj := v.(map[string]any)
		phaseNum := k // phaseNum is a string
		for k, v := range phaseObj {
			if arr, ok := v.([]interface{}); ok {
				fmt.Println(arr)
				// Save phase and store its ID
				phaseID := AddPhase(phaseNum, project.ID, k, int(arr[0].(float64)), int(arr[1].(float64)), int(arr[2].(float64)), arr[3].([]interface{}), arr[4].([]interface{}), arr[5].([]interface{}), arr[6].([]interface{}), arr[7].([]interface{}), arr[8].([]interface{}))
				phaseIDMap[phaseNum] = phaseID
				break // Exit the loop after finding the phase array
			}
		}
	}

	// Step 2: Extract and save all activities
	activityIDMap := make(map[string]map[string]uint)
	for k, v := range projectObj {
		phaseObj := v.(map[string]any)
		phaseNum := k // phaseNum is a string
		phaseID := phaseIDMap[phaseNum]
		activityIDMap[phaseNum] = make(map[string]uint)
		for k, v := range phaseObj {
			if _, ok := v.([]interface{}); ok {
				continue // Skip the phase array as it is already processed
			}
			activityObj := v.(map[string]any)
			actNum := k // actNum is a string
			for k, v := range activityObj {
				if arr, ok := v.([]interface{}); ok {
					// Save activity and store its ID
					activityID := AddActivity(actNum, project.ID, phaseID, k, int(arr[0].(float64)), int(arr[1].(float64)), int(arr[2].(float64)))
					activityIDMap[phaseNum][actNum] = activityID
					break // Exit the loop after finding the activity array
				}
			}
		}
	}

	// Step 3: Iterate again to add sub-activities
	for k, v := range projectObj {
		phaseObj := v.(map[string]any)
		phaseNum := k // phaseNum is a string
		phaseID := phaseIDMap[phaseNum]
		for k, v := range phaseObj {
			if _, ok := v.([]interface{}); ok {
				continue // Skip the phase array as it is already processed
			}
			activityObj := v.(map[string]any)
			actNum := k // actNum is a string
			activityID := activityIDMap[phaseNum][actNum]
			for k, v := range activityObj {
				if _, ok := v.([]interface{}); ok {
					continue // Skip the activity array as it is already processed
				}
				subActivityObj := v.(map[string]any)
				subActNum := k // subActNum is a string
				for k, v := range subActivityObj {
					if arr, ok := v.([]interface{}); ok {
						// Save sub-activity
						AddSubActivity(subActNum, project.ID, activityID, phaseID, k, int(arr[0].(float64)), int(arr[1].(float64)), int(arr[2].(float64)))
					}
				}
			}
		}
	}

	uadmin.ReturnJSON(w, r, map[string]any{
		"status":   "ok",
		"response": project,
	})
}
