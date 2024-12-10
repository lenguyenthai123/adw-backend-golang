package entity

import (
	"time"
)

type Task struct {
	TaskID        int       `gorm:"primaryKey;column:taskId"`
	UserID        int       `gorm:"column:userId"`
	TaskName      string    `gorm:"column:taskName"`
	Description   string    `gorm:"column:description"`
	Priority      string    `gorm:"column:priority;default:Medium"`
	EstimatedTime int       `gorm:"column:estimatedTime"`
	Status        string    `gorm:"column:status;default:Todo"`
	CreatedAt     time.Time `gorm:"column:createdAt;autoCreateTime"`
	DueDate       time.Time `gorm:"column:dueDate"`
	LastUpdated   time.Time `gorm:"column:lastUpdated;autoUpdateTime"`
}

// TableName specifies the custom table name for the Task model (if needed)
func (Task) TableName() string {
	return "Tasks"
}
