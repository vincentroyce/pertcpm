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
	uadmin.StartServer()
}
