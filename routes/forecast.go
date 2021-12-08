package routes

import (
	"github.com/gin-gonic/gin"
	v1Handlers "reperio-backend-assessment/handlers/v1"
)

// forecastRouter is a function for adding the forecast route to a routerGroup
func forecastRouter(group *gin.RouterGroup) (err error) {

	group.GET("/forecast", v1Handlers.GetForecast())

	return
}