package api

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/uadmin/uadmin"
	"github.com/vrsalazar/pertcpm/models"
)

func PredictWeatherAPI(w http.ResponseWriter, r *http.Request) {

	if strings.ToUpper(r.Method) != "GET" {
		// Handle error
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "invalid http method.",
		})
		return
	}

	url := "http://0.0.0.0:5000/generate_predictions"
	// Send the GET request

	file, err := os.Create("weather_ai/predictions_2024.csv")
	if err != nil {
		// Handle error
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "unable to create csv file. " + err.Error(),
		})
		return
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		// Handle error
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "unable to access ai api. " + err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// Create a file to save the CSV data

	// Read the response body
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "unable to write response body to weather csv. " + err.Error(),
		})
		return
	}

	// Initialize model
	weather := []models.Weather{}
	uadmin.All(&weather)

	// Delete each record
	for _, w := range weather {
		uadmin.Delete(&w)
	}

	csvFile, err := os.Open("weather_ai/predictions_2024.csv")
	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "unable to open csv file. " + err.Error(),
		})
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	for {
		// Read each record
		record, err := reader.Read()
		if err == io.EOF {
			// End of file reached
			break
		}
		if err != nil {
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"err_msg": "unable to read records. " + err.Error(),
			})
			return
		}

		saveWeather := models.Weather{}

		// ! Date
		parsedDate, _ := time.Parse("2006-01-02", record[0])
		saveWeather.Date = parsedDate

		// ! Temp
		saveWeather.Temperature, _ = strconv.ParseFloat(record[1], 64)
		saveWeather.Save()

	}

	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"status":   "ok",
		"response": "weather predicted",
	})
}
