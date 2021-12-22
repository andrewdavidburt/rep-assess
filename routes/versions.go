package routes

import (
	"log"
	"reperio-backend-assessment/middleware"

	"github.com/gin-gonic/gin"
)

func buildVersionRouter(server *gin.Engine) error {
	versions := []string{"v1", "v2"}
	for _, version := range versions {
		var (
			routerGroup *gin.RouterGroup
		)
		switch version {
		case "v1":
			routerGroup = server.Group(version, middleware.GenerateRequestID())
		case "v2":
			routerGroup = server.Group(version, middleware.GenerateRequestID(), middleware.IsCurrentLocationPresent())
		default:
			log.Printf("version: %s not found", version)
		}
		//setup current weather route
		if err := currentWeather(routerGroup, version); err != nil {
			log.Panicln(err)
		}
		//setup current forecast router
		if err := forecastRouter(routerGroup); err != nil {
			log.Panicln(err)
		}
	}

	return nil
}
