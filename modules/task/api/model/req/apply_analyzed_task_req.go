package req

type ApplyAnalyzedTaskRequest struct {
	TaskList  []CreateTaskRequest `json:"taskList" binding:"required"`
	StartTime string              `json:"startTime" binding:"required"`
	EndTime   string              `json:"endTime" binding:"required"`
}
