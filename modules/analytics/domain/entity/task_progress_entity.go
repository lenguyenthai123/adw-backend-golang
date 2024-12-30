package entity

type TaskProgressEntity struct {
	TaskID         string `json:"task_id"`
	TaskName       string `json:"task_name"`
	Status         string `json:"status"`
	Description    string `json:"description"`
	TotalSpentTime string `json:"total_spent_time"`
	EstimatedTime  string `json:"estimated_time"`
}
