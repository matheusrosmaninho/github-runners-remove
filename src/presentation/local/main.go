package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/matheusrosmaninho/github-runners-remove/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Starting script to remove workflows ...")

	totalDelete, err := controllers.RemoveWorkflowsAction()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------------------------")
	fmt.Println("Total workflows deleted:", totalDelete)
	fmt.Println("Script finished.")
}
