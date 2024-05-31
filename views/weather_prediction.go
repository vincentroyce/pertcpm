package views

import (
	"net/http"
)

func WeatherPredictionHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	
	// Prepare the data to be sent to the template
	return map[string]interface{}{
		"Title": "Predict Weather",
	}
}
