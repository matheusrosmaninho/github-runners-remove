package entity

import "time"

const (
	RUNNER_STATUS_QUEUED = "queued"
)

type WorkflowRun struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	Conclusion string    `json:"conclusion"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (w *WorkflowRun) IsQueued() bool {
	return w.Status == RUNNER_STATUS_QUEUED
}

func (w *WorkflowRun) IsOldestThan(date time.Time) bool {
	return w.CreatedAt.UTC().Before(date.UTC())
}
