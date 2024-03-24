package main

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
	"github.com/vrsalazar/pertcpm/views"
)

func main() {
	uadmin.RootURL = "/pertcpm-admin/"
	uadmin.Register(
		models.ClientUser{},
	)
	http.HandleFunc("/", uadmin.Handler(views.Main))
	http.HandleFunc("/index/", uadmin.Handler(views.IndexHandler))
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))
	uadmin.StartServer()
}
