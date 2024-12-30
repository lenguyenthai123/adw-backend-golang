package entity

import "time"

type TaskProgress struct {
	ID           int       `gorm:"primaryKey;column:id"`     // Primary key
	UserID       int       `gorm:"column:userId"`            // Foreign key from Users table
	TaskID       int       `gorm:"primaryKey;column:taskId"` // Foreign key from Tasks table
	SessionStart time.Time `gorm:"column:sessionStart"`      // Start time of the session
	SessionEnd   time.Time `gorm:"column:sessionEnd"`        // End time of the session
}

func (TaskProgress) TableName() string {
	return "TimeProgressHistory"
}
