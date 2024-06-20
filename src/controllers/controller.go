package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/matheusrosmaninho/github-runners-remove/helpers"
	"github.com/matheusrosmaninho/github-runners-remove/services"
)

func RemoveWorkflowsAction() error {
	fmt.Println("Starting script to remove workflows ...")
	diasLimite, err := strconv.ParseInt(os.Getenv("INPUT_DAYS_LIMIT"), 10, 32)
	if err != nil {
		message := fmt.Sprintf("Error parsing days limit: %+v", err)
		return fmt.Errorf(message)
	}

	var idsToDelete []int

	workflows, err := services.GetWorkflows(os.Getenv("INPUT_REPO_OWNER"), os.Getenv("INPUT_REPO_NAME"), os.Getenv("INPUT_ACCESS_TOKEN"), 1)
	if err != nil {
		return err
	}

	if workflows.TotalCount == 0 {
		fmt.Println("No workflows found to delete.")
		return nil
	}

	pagination := helpers.NewPagination(workflows.TotalCount, 1)
	for i := 1; i <= pagination.TotalPages; i++ {
		fmt.Println("Getting workflows from page:", i, "of", pagination.TotalPages)
		workflows, err := services.GetWorkflows(os.Getenv("INPUT_REPO_OWNER"), os.Getenv("INPUT_REPO_NAME"), os.Getenv("INPUT_ACCESS_TOKEN"), i)
		if err != nil {
			return err
		}

		for _, workflowRun := range workflows.WorkflowRuns {
			if workflowRun.IsQueued() {
				continue
			}

			actualDate := time.Now()
			dateExpected := helpers.SubtractDaysInDate(actualDate, int(diasLimite))

			if workflowRun.IsOldestThan(dateExpected) {
				idsToDelete = append(idsToDelete, workflowRun.ID)
			}
		}
	}
	fmt.Println("Total workflows to delete in history:", len(idsToDelete))

	for _, id := range idsToDelete {
		err := services.DeleteWorkflow(os.Getenv("INPUT_REPO_OWNER"), os.Getenv("INPUT_REPO_NAME"), os.Getenv("INPUT_ACCESS_TOKEN"), id)
		if err != nil {
			return err
		}
		fmt.Println("Workflow deleted with success:", id)
	}
	return nil
}
