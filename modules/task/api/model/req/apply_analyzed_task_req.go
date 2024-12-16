package req

type ApplyAnalyzedTaskRequest struct {
	TaskList []UpdateTaskRequest `json:"taskList" binding:"required"`
}
