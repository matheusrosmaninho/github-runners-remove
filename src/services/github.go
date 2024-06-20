package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/matheusrosmaninho/github-runners-remove/entity"
	"github.com/matheusrosmaninho/github-runners-remove/helpers"
)

const (
	GITHUB_BASE_URL = "https://api.github.com"
	GITHUB_REPO_URL = GITHUB_BASE_URL + "/repos"
)

type ResponseWorkflowsList struct {
	TotalCount   int                  `json:"total_count"`
	WorkflowRuns []entity.WorkflowRun `json:"workflow_runs"`
}

var headers = map[string]string{
	"Accept":               "application/vnd.github+json",
	"X-GitHub-Api-Version": "2022-11-28",
}

func GetWorkflows(owner, repo, accessToken string, page int) (*ResponseWorkflowsList, error) {
	url := fmt.Sprintf("%s/%s/%s/actions/runs?page=%d&per_page=%d", GITHUB_REPO_URL, owner, repo, page, helpers.ITEMS_PER_PAGE)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		message := fmt.Sprintf("Error creating request: %+v", err)
		return nil, fmt.Errorf(message)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := client.Do(req)
	if err != nil {
		message := fmt.Sprintf("Error making request: %+v", err)
		return nil, fmt.Errorf(message)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		message := fmt.Sprintf("Error reading response: %+v", err)
		return nil, fmt.Errorf(message)
	}

	var response ResponseWorkflowsList
	err = json.Unmarshal(body, &response)
	if err != nil {
		message := fmt.Sprintf("Error decoding response: %+v", err)
		return nil, fmt.Errorf(message)
	}
	return &response, nil
}

func DeleteWorkflow(owner, repo, accessToken string, id int) error {
	url := fmt.Sprintf("%s/%s/%s/actions/runs/%d", GITHUB_REPO_URL, owner, repo, id)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		message := fmt.Sprintf("Error creating request: %+v", err)
		return fmt.Errorf(message)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := client.Do(req)
	if err != nil {
		message := fmt.Sprintf("Error making request: %+v", err)
		return fmt.Errorf(message)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		message := fmt.Sprintf("Error deleting workflow: %+v", resp.Status)
		return fmt.Errorf(message)
	}
	return nil
}
