package packages

import (
	"fmt"
	"net/http"
	"reperio-backend-assessment/errors"
	"reperio-backend-assessment/models"
	"reperio-backend-assessment/types"
)


var (
	// WeatherApi is a variable pointed to an interface so that we can make requests to the api
	WeatherApi types.WeatherApi = &weather{apiKey: "4234d954d8024273b7925534210712", host: "api.weatherapi.com"}
)

type weather struct {
	apiKey string
	host   string
}

func (w *weather) getWrapper(url string, params ...interface{}) (resp *http.Response, err error) {
	resp, err = get(fmt.Sprintf(url, params...))
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("Response for %s => GET[Status: %d]", url, resp.StatusCode))
	return
}

// Forecast is a public method for fetching the current forecast for a location based on a number of days
func (w *weather) Forecast(location, days string) (record *models.ForecastRecord, err error) {
	resp, err := w.getWrapper("https://%s/v1/forecast.json?key=%s&q=%s&days=%d", w.host, w.apiKey, location, days)
	
	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		record = &models.ForecastRecord{}
		if err = parseBody(resp, record); err != nil {
			return
		}
		return
	}

	return nil, errors.NewHTTPError(resp.StatusCode, nil)
}

// Current is a public method for fetching the current weather from the api
func (w *weather) Current(location string) (record *models.CurrentWeather, err error) {
	resp, err := get(fmt.Sprintf("https://%s/v1/current.json?key=&q=%s&aqi=no&key=%s", w.host, location,w.apiKey))
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusOK {
		record = &models.CurrentWeather{}
		if err = parseBody(resp, record); err != nil {
			return
		}
		return
	}

	return nil, errors.NewHTTPError(resp.StatusCode, nil)
}
