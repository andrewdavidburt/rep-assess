package types

import "reperio-backend-assessment/models"

// WeatherApi is an interface detailing what methods the weatherapi package should provide
// helps with mocking libraries
type WeatherApi interface {
	Current(location string) (record *models.CurrentWeather, err error)
	Forecast(location string, days string) (record *models.ForecastRecord, err error)
}
