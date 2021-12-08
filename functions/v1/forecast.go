package v1functions

import (
	"reperio-backend-assessment/models"
	"reperio-backend-assessment/packages"
)

// ForecastWeather is a standalone function that interfaces with sdk. This was moved here so that if we decided to go
// serverless the raw function logic was reusable with minimal changes
func ForecastWeather(location, days string) (currentWeatherRecord *models.ForecastRecord, err error) {
	currentWeatherRecord, err = packages.WeatherApi.Forecast(location, days)
	return
}