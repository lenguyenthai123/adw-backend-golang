package req

type ListIDRequest struct {
	TaskIDList []string `json:"taskList" binding:"required"`
}
