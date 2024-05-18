package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func CompletedProjectHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/ongoing-projects/")

	projects := []models.Project{}
	uadmin.Filter(&projects, "completed = ?", true)
	fmt.Println(projects)

	return map[string]interface{}{
		"Title":    "Completed Projects",
		"Projects": projects,
	}
}
