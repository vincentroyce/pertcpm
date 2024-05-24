package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/dashboard/")

	allProject := models.Project{}

	allProjectCount := uadmin.Count(&allProject, "id > 0")
	ongoingProjectCount := uadmin.Count(&allProject, "completed = ?", false)
	completedProjectCount := uadmin.Count(&allProject, "completed = ?", true)

	return map[string]interface{}{
		"Title":             "User Dashboard",
		"AllProjects":       allProjectCount,
		"OngoingProjects":   ongoingProjectCount,
		"CompletedProjects": completedProjectCount,
	}
}
