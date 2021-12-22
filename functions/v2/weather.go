package v1functions

import (
	"fmt"
	"reperio-backend-assessment/database"
	"reperio-backend-assessment/models"
	"reperio-backend-assessment/packages"
)

// CurrentWeather is a standalone function that interfaces with sdk. This was moved here so that if we decided to go
// serverless the raw function logic was reusable with minimal changes
func CurrentWeather(location interface{}) (currentWeatherRecord *models.CurrentWeather, err error) {

	var loc string = location.(string)

	currentWeatherRecord, err = packages.WeatherApi.Current(loc)

	record := &models.CurrentWeatherDatabase{
		Location:    loc,
		TempF:       currentWeatherRecord.Current.TempF,
		TempC:       currentWeatherRecord.Current.TempC,
		LastUpdated: currentWeatherRecord.Current.LastUpdatedEpoch,
	}

	success, err := database.Insert("current_weather", record)
	fmt.Println(success, err)

	return
}
