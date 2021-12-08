package v1Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reperio-backend-assessment/errors"
	v1functions "reperio-backend-assessment/functions/v1"
)

// GetForecast is a handler function for interacting with the sdk
func GetForecast() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			location = context.Query("location")
			days  = context.Query("days")
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
		context.AbortWithStatusJSON(http.StatusOK, res)
	}
}