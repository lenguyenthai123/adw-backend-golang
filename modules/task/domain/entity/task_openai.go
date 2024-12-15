package entity

// TaskOpenai struct đại diện cho một nhiệm vụ
type TaskOpenai struct {
	TaskID        int    `json:"taskId" jsonschema_description:"The unique ID of the task"`
	UserID        int    `json:"userId" jsonschema_description:"The ID of the user who owns the task"`
	TaskName      string `json:"taskName" jsonschema_description:"The name of the task"`
	Description   string `json:"description" jsonschema_description:"The detailed description of the task"`
	Priority      string `json:"priority" jsonschema:"enum=Low,enum=Medium,enum=High" jsonschema_description:"The priority level of the task"`
	EstimatedTime int    `json:"estimatedTime" jsonschema_description:"Estimated time in hours to complete the task"`
	Status        string `json:"status" jsonschema:"enum=Todo,enum=InProgress,enum=Completed,enum=Expired" jsonschema_description:"The current status of the task"`
	DueDate       string `json:"dueDate" jsonschema_description:"Deadline for the task"`
	LastUpdated   string `json:"lastUpdated" jsonschema_description:"Timestamp when the task was last updated"`
}
