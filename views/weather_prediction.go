package views

import (
	"net/http"
	"time"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func WeatherPredictionHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	weather := []models.Weather{}
	uadmin.Filter(&weather, "date > ?", time.Now().Add(-24*time.Hour))

	weather_view := []map[string]interface{}{}

	for i := range weather {
		weather_view = append(weather_view, map[string]interface{}{
			"ID":          weather[i].ID,
			"Date":        weather[i].Date.Format("January 2, 2006"),
			"Temperature": weather[i].Temperature,
			"Status":      weather[i].Status,
		})
	}

	// Prepare the data to be sent to the template
	return map[string]interface{}{
		"Title":      "Predict Weather",
		"Prediction": weather_view,
	}
}
