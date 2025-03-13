package main

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/api"
	"github.com/vrsalazar/pertcpm/models"
	"github.com/vrsalazar/pertcpm/views"
)

func main() {
	uadmin.RootURL = "/pertcpm-admin/"
	uadmin.Register(
		models.ClientUser{},
		models.Project{},
		models.Equipment{},
		models.Worker{},
		models.Phase{},
		models.Activity{},
		models.SubActivity{},
		models.Weather{},
	)

	uadmin.RegisterInlines(
		models.Project{},
		map[string]string{
			"Equipment":   "ProjectID",
			"Worker":      "ProjectID",
			"Phase":       "ProjectID",
			"Activity":    "ProjectID",
			"SubActivity": "ProjectID",
		},
	)
	uadmin.Port = 1111
	http.HandleFunc("/", uadmin.Handler(views.Main))
	http.HandleFunc("/api/", uadmin.Handler(api.Main))
	http.HandleFunc("/index/", uadmin.Handler(views.IndexHandler))
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))
	CustomizeModel()
	uadmin.StartServer()
}

func CustomizeModel() {
	modelschema := uadmin.ModelSchema{
		Name:          "Project", // Model name
		ModelName:     "project", // URL
		IncludeListJS: []string{"https://cdn.jsdelivr.net/npm/sweetalert2@11", "/static/js/list.js"},
	}

	// Call the schema of "todo" model
	// modelschema.ModelName = "todo"
	project := uadmin.Schema[modelschema.ModelName]
	// Include Javascript file for the list
	project.IncludeListJS = modelschema.IncludeListJS
	uadmin.Schema[modelschema.ModelName] = project
}
