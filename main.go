package main

import (
	"fmt"
	"log"
	"net/http"
	"reperio-backend-assessment/database"
	"reperio-backend-assessment/env"
	"reperio-backend-assessment/routes"
)

func main () {
	env.Setup()

	if err := database.Setup(); err != nil {
		log.Panicln(err)
	}

	server := routes.BuildEngine()

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", env.PORT), server))
}
