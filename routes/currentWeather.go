package routes

import (
	v1Handlers "reperio-backend-assessment/handlers/v1"

	v2Handlers "reperio-backend-assessment/handlers/v2"

	"github.com/gin-gonic/gin"
)

func currentWeather(group *gin.RouterGroup, version string) error {
	switch version {
	case "v1":
		group.GET("current-weather", v1Handlers.CurrentWeather())
	case "v2":
		group.GET("current-weather", v2Handlers.CurrentWeather())
	}
	return nil
}
