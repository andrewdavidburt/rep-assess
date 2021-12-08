package env

import (
	"log"
	"os"
)

var (
	PORT, NAME string
)

func Setup() {
	PORT = os.Getenv("PORT")
	NAME = os.Getenv("ENVIRONMENT")

	if PORT == "" {
		log.Fatalln("PORT MUST BE SET")
	}

	if NAME == "" {
		log.Fatalln("ENVIRONMENT MUST BE SET")
	}
}