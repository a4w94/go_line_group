package main

import (
	"go_line_group/internal/route"
	"log"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := route.InitRouter()

	router.Run(":" + port)
}
