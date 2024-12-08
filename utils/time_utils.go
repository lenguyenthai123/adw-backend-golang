package utils

import (
	"backend-golang/pkgs/log"
	"log/slog"
	"time"
)

func ResetTime() *string {
	now := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	return ParseTimeToSearchSQL(now)
}

func ParseTimeToSearchSQL(time time.Time) *string {
	timeFormat := time.Format("2006-01-02")
	return &timeFormat
}

func ParseTime(layout, inputTime string) *time.Time {
	result, err := time.Parse(layout, inputTime)
	if err != nil {
		log.JsonLogger.Error("Error ParseTime",
			slog.String("layout", layout),
			slog.String("inputTime", inputTime),
			slog.String("error", err.Error()),
		)
		return nil
	}
	return &result
}

func Add1DayInputString(input string) *string {
	var timeTo = ParseTime("2006-01-02", input)
	return ParseTimeToSearchSQL(timeTo.AddDate(0, 0, 1))
}

func Add1DayInputTime(input time.Time) *string {
	return ParseTimeToSearchSQL(input.AddDate(0, 0, 1))
}
