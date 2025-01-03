package entity

import "time"

type TaskOpenai struct {
	TaskName    string    `json:"taskName"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	DueDate     time.Time `json:"dueDate"`
	Priority    string    `json:"priority"`
	Status      string    `json:"status"`
}

type Feedback struct {
	Strengths         []string `json:"strengths"`
	AttentionAreas    []string `json:"attentionAreas"`
	MotivationalNotes []string `json:"motivationalNotes"`
}

type TaskOpenaiResponse struct {
	TaskList []TaskOpenai `json:"taskList"`
	Feedback Feedback     `json:"feedback"`
}
