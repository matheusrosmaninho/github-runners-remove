package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/matheusrosmaninho/github-runners-remove/helpers"
	"github.com/matheusrosmaninho/github-runners-remove/services"
)

func RemoveWorkflowsAction() (int, error) {
	diasLimite, err := strconv.ParseInt(os.Getenv("INPUT_DAYS_LIMIT"), 10, 32)
	if err != nil {
		message := fmt.Sprintf("Error parsing days limit: %+v", err)
		return 0, fmt.Errorf(message)
	}

	var idsToDelete []int

	workflows, err := services.GetWorkflows(os.Getenv("INPUT_REPO_OWNER"), os.Getenv("INPUT_REPO_NAME"), os.Getenv("INPUT_ACCESS_TOKEN"), 1)
	if err != nil {
		return 0, err
	}

	if workflows.TotalCount == 0 {
		fmt.Println("No workflows found to delete.")
		return 0, nil
	}

	pagination := helpers.NewPagination(workflows.TotalCount, 1)
	for i := 1; i <= pagination.TotalPages; i++ {
		fmt.Println("Getting workflows from page:", i, "of", pagination.TotalPages)
		workflows, err := services.GetWorkflows(os.Getenv("INPUT_REPO_OWNER"), os.Getenv("INPUT_REPO_NAME"), os.Getenv("INPUT_ACCESS_TOKEN"), i)
		if err != nil {
			return 0, err
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

	for index, id := range idsToDelete {
		err := services.DeleteWorkflow(os.Getenv("INPUT_REPO_OWNER"), os.Getenv("INPUT_REPO_NAME"), os.Getenv("INPUT_ACCESS_TOKEN"), id)
		if err != nil {
			return 0, err
		}
		fmt.Println("Workflow deleted with success:", id, ". Remaining:", len(idsToDelete)-index-1)
	}
	return len(idsToDelete), nil
}
