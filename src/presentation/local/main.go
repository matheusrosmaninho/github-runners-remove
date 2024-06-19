package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/matheusrosmaninho/github-runners-remove/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = controllers.RemoveWorkflowsAction()
	if err != nil {
		log.Fatal(err)
	}
}
