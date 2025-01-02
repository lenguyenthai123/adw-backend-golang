package req

type CreateTaskRequest struct {
	TaskName      string `json:"taskName" binding:"required"`
	Description   string `json:"description"`
	Priority      string `json:"priority" binding:"required,oneof=Low Medium High"`
	EstimatedTime string `json:"estimatedTime"`
	Status        string `json:"status" binding:"required,oneof=Todo InProgress Completed Expired"`
	StartDate     string `json:"startDate"` // ISO format date string
	DueDate       string `json:"dueDate"`   // ISO format date string
}
