package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

func BuildEngine() *gin.Engine {
	server := gin.New()

	if err := buildVersionRouter(server); err != nil {
		log.Panic(err)
	}

	return server
}