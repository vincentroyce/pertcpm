package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func Main(w http.ResponseWriter, r *http.Request) {
	page := strings.TrimPrefix(r.URL.Path, "/user/")
	page = strings.TrimSuffix(page, "/")
	context := map[string]interface{}{}
	sess := uadmin.IsAuthenticated(r)

	if sess == nil {
		http.Redirect(w, r, "/index/", http.StatusSeeOther)
	}

	switch page {
	case "dashboard":
		context = DashboardHandler(w, r)
	case "home":
		context = HomeHandler(w, r)
	case "pert":
		context = PertHandler(w, r)
	case "planning":
		context = PlanningHandler(w, r)
	case "schedule-completion":
		context = ScheduleCompletionHandler(w, r)
	case "ongoing-projects":
		context = OngoingProjectHandler(w, r)
	default:
		page = "home"
	}
	context["Page"] = page
	Render(w, r, page, context)
}

func Render(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	tmpList := []string{}
	tmpList = append(tmpList, "./templates/base.html")
	path := "./templates/" + page + ".html"
	tmpList = append(tmpList, path)

	uadmin.RenderMultiHTML(w, r, tmpList, context)
}

func PertHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/pert/")
	return map[string]interface{}{
		"Title": "PERT",
	}
}

func PlanningHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/planning/")
	return map[string]interface{}{
		"Title": "Planning",
	}
}
