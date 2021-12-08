package routes

import (
	"github.com/gin-gonic/gin"
	v1Handlers "reperio-backend-assessment/handlers/v1"
)

func currentWeather(group *gin.RouterGroup) error {
	group.GET("current-weather", v1Handlers.CurrentWeather())
	return nil
}