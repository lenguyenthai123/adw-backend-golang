package entity

import (
	"time"
)

type TaskApplyAnalyzedTaskEntity struct {
	TaskList  []*Task   `json:"taskList" binding:"required"`
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
}
