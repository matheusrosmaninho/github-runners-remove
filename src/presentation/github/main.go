package main

import (
	"fmt"
	"log"

	"github.com/matheusrosmaninho/github-runners-remove/controllers"
)

func main() {
	fmt.Println("Starting script to remove workflows ...")

	totalDelete, err := controllers.RemoveWorkflowsAction()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------------------------")
	fmt.Println("Total workflows deleted:", totalDelete)
	fmt.Println("Script finished.")
}
