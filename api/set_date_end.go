package api

import (
	"fmt"
	"net/http"

	"github.com/uadmin/uadmin"
)

func SetDateEndAPI(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"err_msg": "unable to parse form. " + err.Error(),
			"status":  "error",
		})
		return
	}

	id := r.FormValue("id")
	days := r.FormValue("days")

	fmt.Println(id)
	fmt.Println(days)

}
