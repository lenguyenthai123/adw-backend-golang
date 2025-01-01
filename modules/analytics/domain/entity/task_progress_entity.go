package entity

type TaskProgressEntity struct {
	TaskID         string `gorm:"column:taskId"`
	TaskName       string `gorm:"column:taskName"`
	Status         string `gorm:"column:status"`
	TotalSpentTime string `gorm:"column:totalTimeSpent"`
	EstimatedTime  string `gorm:"column:estimatedTime"`
}
