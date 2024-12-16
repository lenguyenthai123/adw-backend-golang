package req

type UpdateTaskRequest struct {
	TaskID        string `json:"taskId"`
	TaskName      string `json:"taskName"`
	Description   string `json:"description"`
	Priority      string `json:"priority" binding:"oneof=Low Medium High"`
	EstimatedTime int    `json:"estimatedTime"`
	Status        string `json:"status" binding:"oneof=Todo InProgress Completed Expired"`
	DueDate       string `json:"dueDate"`   // ISO format date string
	StartDate     string `json:"startDate"` // ISO format date string
}
