package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func ScheduleCompletionHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// Parse project ID from the request URL
	// projectID := r.URL.Query().Get("1")
	// if projectID == "" {
	// 	http.Error(w, "Project ID is required", http.StatusBadRequest)
	// }

	// Fetch the project from the database
	project := models.Project{}
	err := uadmin.Get(&project, "id = ?", 1)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
	}
	name := project.Name
	// Load related phases for the project
	phases := []models.Phase{}
	uadmin.Filter(&phases, "project_id = ?", project.ID)
	// Load related activities and sub-activities for each phase
	for i := range phases {
		activities := []models.Activity{}
		uadmin.Filter(&activities, "phase_id = ?", phases[i].ID)
		phases[i].Activity = activities
		for j := range phases[i].Activity {
			subActivities := []models.SubActivity{}
			uadmin.Filter(&subActivities, "activity_id = ?", phases[i].Activity[j].ID)
			phases[i].Activity[j].SubActivity = subActivities
		}
	}

	return map[string]interface{}{
		"Title":         "Schedule Completion",
		"ProjectName":   name,
		"Project":       project,
		"Phases":        phases,
		"Activities":    activities,
		"SubActivities": subActivities,
	}
}
