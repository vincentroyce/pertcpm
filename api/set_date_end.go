package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
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
	
	project := models.Project{}
	err = uadmin.Get(&project, "id = ?", id)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"err_msg": "unable to get the project. " + err.Error(),
			"status":  "error",
		})
		return
	}

	addDays, err := strconv.Atoi(days)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"err_msg": "unable to convert string days to int. " + err.Error(),
			"status":  "error",
		})
		return
	}

	startDate := project.DateStart
	dateEnd := startDate.AddDate(0, 0, addDays)

	project.DateEnd = &dateEnd
	project.Save()

	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"response": "Date end have been updated",
		"status":   "ok",
	})
}
