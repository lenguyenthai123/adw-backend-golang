package req

type CreateTaskRequest struct {
	UserID        int    `json:"userId" binding:"required"`
	TaskName      string `json:"taskName" binding:"required"`
	Description   string `json:"description"`
	Priority      string `json:"priority" binding:"required,oneof=Low Medium High"`
	EstimatedTime int    `json:"estimatedTime"`
	Status        string `json:"status" binding:"required,oneof=Todo InProgress Completed Expired"`
	DueDate       string `json:"dueDate"` // ISO format date string
}
