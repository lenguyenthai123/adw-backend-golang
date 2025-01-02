package res

import "time"

type TaskResponse struct {
	TaskID        int       `json:"taskId"`
	UserID        int       `json:"userId"`
	TaskName      string    `json:"taskName"`
	Description   string    `json:"description"`
	Priority      string    `json:"priority"`
	EstimatedTime string    `json:"estimatedTime"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	StartDate     time.Time `json:"startDate"`
	DueDate       time.Time `json:"dueDate"`
	LastUpdated   time.Time `json:"lastUpdated"`
}
