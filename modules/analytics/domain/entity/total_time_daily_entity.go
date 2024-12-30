package entity

type DailyProgressEntity struct {
	Date      string `gorm:"column:day"`
	TotalTime string `json:"spent_time"` // Total hours spent at that day
}
