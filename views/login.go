package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login/")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	context := map[string]interface{}{}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		session, _ := uadmin.Login(r, username, password)
		if session != nil {
			http.SetCookie(w, &http.Cookie{
				Name:     "session",
				Value:    session.Key,
				Path:     "/",
				SameSite: http.SameSiteStrictMode,
			})
			http.Redirect(w, r, "/user/dashboard/", http.StatusSeeOther)
		} else {
			context["Invalid"] = "Please put a valid credentials."
		}
	}
	uadmin.RenderHTML(w, r, "templates/login.html", context)
}
