package v1Handlers

import (
	"net/http"
	"reperio-backend-assessment/errors"
	v2functions "reperio-backend-assessment/functions/v2"

	"github.com/gin-gonic/gin"
)

// CurrentWeather is a handler used with Gin. This can easily be replaced with something like mux router or gorilla or
// martini
func CurrentWeather() gin.HandlerFunc {
	return func(context *gin.Context) {
		location := context.Query("location")
		if location == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "location is required in url query",
			})
			return
		}

		data, err := v2functions.CurrentWeather(location)
		if err != nil {
			if httpError, ok := err.(*errors.HTTPError); ok {
				context.AbortWithStatusJSON(httpError.StatusCode, httpError.Message)
				return
			}
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		context.AbortWithStatusJSON(http.StatusOK, data)
	}
}
