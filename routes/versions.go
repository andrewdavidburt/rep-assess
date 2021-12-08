package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"reperio-backend-assessment/middleware"
)

func buildVersionRouter(server *gin.Engine) error {
	versions := []string{"v1"}
	for _, version := range versions {
		var (
			routerGroup *gin.RouterGroup
		)
		switch version {
		case "v1":
			routerGroup = server.Group(version, middleware.GenerateRequestID())
		default:
			log.Printf("version: %s not found", version)
		}
		//setup current weather route
		if err := currentWeather(routerGroup); err != nil {
			log.Panicln(err)
		}
		//setup current forecast router
		if err := forecastRouter(routerGroup); err != nil {
			log.Panicln(err)
		}
	}

	return nil
}