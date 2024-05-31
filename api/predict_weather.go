package api

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/uadmin/uadmin"
)

func PredictWeatherAPI(w http.ResponseWriter, r *http.Request) {
	// Run python
	cmd := exec.Command("python3", "weather_ai/api.py")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	host, err := os.Hostname()
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "unable to get hostname. " + err.Error(),
		})
		return
	}

	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"status":   "ok",
		"response": host,
	})
}
