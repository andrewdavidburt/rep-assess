package v1Handlers

import (
	"net/http"
	"reperio-backend-assessment/errors"
	v1functions "reperio-backend-assessment/functions/v1"
	"sort"

	"github.com/gin-gonic/gin"
)

// GetForecast is a handler function for interacting with the sdk
func GetForecast() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			location = context.Query("location")
			days     = context.Query("days")
			sortdir  = context.Query("sort")
		)

		res, err := v1functions.ForecastWeather(location, days)
		if err != nil {
			if httpError, ok := err.(*errors.HTTPError); ok {
				context.AbortWithStatusJSON(httpError.StatusCode, httpError.Message)
				return
			}
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		if sortdir == "asc" {
			sort.Slice(res.Forecast.Forecastday, func(i, j int) bool {
				return res.Forecast.Forecastday[i].Day.AvgtempC < res.Forecast.Forecastday[j].Day.AvgtempC
			})

		}

		if sortdir == "desc" {
			sort.Slice(res.Forecast.Forecastday, func(i, j int) bool {
				return res.Forecast.Forecastday[i].Day.AvgtempC > res.Forecast.Forecastday[j].Day.AvgtempC
			})

		}
		context.AbortWithStatusJSON(http.StatusOK, res)
	}
}
