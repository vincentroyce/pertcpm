package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func EditPredecessorsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	projectID := r.URL.Query().Get("id")
	if projectID == "" {
		http.Error(w, "Project ID is required", http.StatusBadRequest)
	}

	// Fetch the project from the database
	project := models.Project{}
	err := uadmin.Get(&project, "id = ?", projectID)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
	}

	// Load related phases for the project
	phases := []models.Phase{}
	uadmin.Filter(&phases, "project_id = ?", project.ID)

	allSubActivities := []models.SubActivity{}
	// Loop through each phase to load related activities and sub-activities
	for i := range phases {
		// Load activities for the phase
		activities := []models.Activity{}
		uadmin.Filter(&activities, "phase_id = ?", phases[i].ID)

		// Loop through each activity to load related sub-activities
		for j := range activities {
			subActivities := []models.SubActivity{}
			uadmin.Filter(&subActivities, "activity_id = ?", activities[j].ID)
			activities[j].SubActivity = subActivities
			allSubActivities = append(allSubActivities, subActivities...)
		}

		// Assign activities to the phase
		phases[i].Activity = activities
	}

	// Assign phases to the project
	project.Phase = phases

	// Prepare the data to be sent to the template
	return map[string]interface{}{
		"Title":            "View Project",
		"ProjectName":      project.Name,
		"Project":          project,
		"AllSubActivities": allSubActivities,
	}
}
