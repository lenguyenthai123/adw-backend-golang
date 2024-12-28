package req

import (
	"fmt"
	"time"
)

type UpdateTaskProgressRequest struct {
	SessionStart string `json:"sessionStart" binding:"required"`
	SessionEnd   string `json:"sessionEnd" binding:"required"`
}

func (r *UpdateTaskProgressRequest) Validate() error {
	// Parse the sessionStart and sessionEnd strings into time.Time objects
	layout := time.RFC3339 // Use RFC3339 format for time parsing
	startTime, err := time.Parse(layout, r.SessionStart)
	if err != nil {
		return fmt.Errorf("invalid sessionStart format: %v", err)
	}

	endTime, err := time.Parse(layout, r.SessionEnd)
	if err != nil {
		return fmt.Errorf("invalid sessionEnd format: %v", err)
	}

	// Validate that sessionEnd is after sessionStart
	if !endTime.After(startTime) {
		return fmt.Errorf("sessionEnd must be after sessionStart")
	}

	return nil
}
